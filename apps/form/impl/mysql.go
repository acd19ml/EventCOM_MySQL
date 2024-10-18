package impl

import (
	"database/sql"

	"github.com/acd19ml/EventCOM_MySQL/apps/form"
	"github.com/acd19ml/EventCOM_MySQL/conf"
	"github.com/acd19ml/EventCOM_MySQL/mcube/logger"
	"github.com/acd19ml/EventCOM_MySQL/mcube/logger/zap"
)

type FormServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}

// 接口实现静态检查
var _ form.Service = (*FormServiceImpl)(nil)

// 保证调用该函数之前, 全局conf对象已经初始化
func NewFormServiceImpl() *FormServiceImpl {
	return &FormServiceImpl{
		// Form service 服务的子Loggger
		// 封装的Zap让其满足 Logger接口
		// 为什么要封装:
		// 		1. Logger全局实例
		// 		2. Logger Level的动态调整, Logrus不支持Level共同调整
		// 		3. 加入日志轮转功能的集合
		l:  zap.L().Named("Form"),
		db: conf.C().MySQL.GetDB(),
	}
}
