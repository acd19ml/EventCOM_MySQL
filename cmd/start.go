package cmd

import (
	"github.com/acd19ml/EventCOM_MySQL/apps/form/http"
	"github.com/acd19ml/EventCOM_MySQL/apps/form/impl"
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

		// 加载我们Host Service的实体类
		// host service 的具体实现
		service := impl.NewFormServiceImpl()

		// 通过Form API Handler 提供 Restful API
		api := http.NewFormHTTPHandler(service)

		// 提供gin router
		g := gin.Default()
		api.Registry(g)

		// 3. 启动服务
		return g.Run(conf.C().App.HTTPAddr())
	},
}

// 还没有初始化Logger实例
func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "EventCOM api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
