package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/acd19ml/EventCOM_MySQL/apps/form"
	"github.com/acd19ml/EventCOM_MySQL/apps/form/impl"
	"github.com/acd19ml/EventCOM_MySQL/mcube/logger/zap"
)

var (
	// 定义对象是满足该接口的实例
	service form.Service
)

func TestCreate(t *testing.T) {
	// 创建一个Form对象
	ins := form.NewForm()
	ins.Name = "test"
	// 调用CreateForm方法
	service.CreateForm(context.Background(), ins)
}

func init() {
	//需要初始化全局logger
	//为什么不设计为默认打印，因为性能
	fmt.Println(zap.DevelopmentSetup())

	// form service 的具体实现
	service = impl.NewFormServiceImpl()
}
