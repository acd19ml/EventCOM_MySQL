package cmd

import (
	"github.com/acd19ml/EventCOM_MySQL/apps"
	_ "github.com/acd19ml/EventCOM_MySQL/apps/all"
	"github.com/acd19ml/EventCOM_MySQL/conf"
	"github.com/gin-gonic/gin"
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
			panic(err)
		}

		// 初始化所有的服务
		apps.InitImpl()

		// 提供gin router
		g := gin.Default()
		// 注册IOC的所有http handler
		apps.InitGin(g)

		// 3. 启动服务
		return g.Run(conf.C().App.HTTPAddr())
	},
}

// 还没有初始化Logger实例
func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "EventCOM api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
