package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/acd19ml/EventCOM_MySQL/apps"
	_ "github.com/acd19ml/EventCOM_MySQL/apps/all"
	"github.com/acd19ml/EventCOM_MySQL/conf"
	"github.com/acd19ml/EventCOM_MySQL/mcube/logger"
	"github.com/acd19ml/EventCOM_MySQL/mcube/logger/zap"
	"github.com/acd19ml/EventCOM_MySQL/protocol"
	"github.com/spf13/cobra"
)

var (
	// pusher service config options
	confType string
	confFile string
	confETCD string
)

// 程序启动组装
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动EventCOM 后端API",
	Long:  "启动EventCOM 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载配置
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			return (err)
		}
		//初始化Logger
		if err := loadGlobalLogger(); err != nil {
			return err
		}

		// 初始化所有的服务
		apps.InitImpl()

		svc := newManager()

		ch := make(chan os.Signal, 1)
		// 中断信号处理
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
		// 协程服务不断监听信号
		go svc.WaitStop(ch)

		// 3. 启动服务
		return svc.Start()
	},
}

func newManager() *manager {
	return &manager{
		http: protocol.NewHttpService(),
		l:    zap.L().Named("CLI"),
	}
}

// 用于管理所有需要启动的服务
// 1. HTTP服务的启动
type manager struct {
	http *protocol.HttpService
	l    logger.Logger
}

func (m *manager) Start() error {
	// 启动HTTP服务
	if err := m.http.Start(); err != nil {
		return err
	}
	return nil
}

// 处理来自外部的中断信号，Terminal信号
func (m *manager) Stop() error {
	return nil
}

// 处理来自外部的中断信号, 比如Terminal
func (m *manager) WaitStop(ch <-chan os.Signal) { //只读channel
	for v := range ch {
		switch v {
		default:
			m.l.Infof("receive signal: %s", v)
			m.http.Stop()
		}
	}

}

// 问题:
// 1. http API, Grpc API 需要启动, 消息总线也需要监听, 比如负责注册于配置,  这些模块都是独立
//    都需要在程序启动时，进行启动, 都写在start start膨胀到不易维护
// 2. 服务的优雅关闭怎么办? 外部都会发送一个Terminal(中断)信号给程序, 程序时需要处理这个信号
//    需要实现程序优雅关闭的逻辑的处理: 由先后顺序 (从外到内完成资源的释放逻辑处理)
//    1. api 层的关闭 (HTTP, GRPC)
//    2. 消息总线关闭
//    3. 关闭数据库链接
//    4. 如果 使用了注册中心, 完成注册中心的注销操作
//    5. 退出完毕

// log 为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化
func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)

	// 更加Config里面的日志配置，来配置全局Logger对象
	lc := conf.C().Log

	// 解析日志Level配置
	// DebugLevel: "debug",
	// InfoLevel:  "info",
	// WarnLevel:  "warning",
	// ErrorLevel: "error",
	// FatalLevel: "fatal",
	// PanicLevel: "panic",
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}

	// 使用默认配置初始化Logger的全局配置
	zapConfig := zap.DefaultConfig()

	// 配置日志的Level基本
	zapConfig.Level = level

	// 程序没启动一次, 不必都生成一个新日志文件
	zapConfig.Files.RotateOnStartup = false

	// 配置日志的输出方式
	switch lc.To {
	case conf.ToStdout:
		// 把日志打印到标准输出
		zapConfig.ToStderr = true
		// 并没在把日志输入输出到文件
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}

	// 配置日志的输出格式:
	switch lc.Format {
	case conf.JSONFormat:
		zapConfig.JSON = true
	}

	// 把配置应用到全局Logger
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}

	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "EventCOM api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
