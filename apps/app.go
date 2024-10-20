package apps

import (
	"fmt"

	"github.com/acd19ml/EventCOM_MySQL/apps/form"
	"github.com/gin-gonic/gin"
)

// IOC 容器层: 管理所有的服务的实例

// 1. FormService的实例必须注册过来, FormService才会有具体的实例, 服务启动时注册
// 2. HTTP 暴露模块, 依赖Ioc里面的FormService
var (
	FormService form.Service
	// 维护当前所有服务
	// svcs = map[string]Service{}
	implApps = map[string]ImplService{}
	ginApps  = map[string]GinService{}
)

type ImplService interface {
	Config()
	Name() string
}

func RegistryImpl(svc ImplService) {
	// 检查服务是否已经注册过
	if _, ok := implApps[svc.Name()]; ok {
		// 如果该服务已经注册，抛出一个 panic 错误，避免重复注册
		panic(fmt.Sprintf("service %s already registered", svc.Name()))
	}

	// 将服务注册到 svcs 容器，键是服务的名字，值是该服务的实例
	implApps[svc.Name()] = svc
	// 判断传入的服务是否满足 form.Service 接口
	if v, ok := svc.(form.Service); ok {
		// 如果是 form.Service 类型的服务，将其赋值给全局的 FormService
		FormService = v
	}
}

// Get 一个Impl服务的实例：implApps
// 返回一个对象, 任何类型都可以, 使用时, 由使用方进行断言
func GetImpl(name string) interface{} {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}

	return nil
}

type GinService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}

func RegistryGin(svc GinService) {
	// 检查服务是否已经注册过
	if _, ok := ginApps[svc.Name()]; ok {
		// 如果该服务已经注册，抛出一个 panic 错误，避免重复注册
		panic(fmt.Sprintf("service %s already registered", svc.Name()))
	}

	// 将服务注册到 svcs 容器，键是服务的名字，值是该服务的实例
	ginApps[svc.Name()] = svc
}

func InitGin(r gin.IRouter) {

	// 初始化对象
	for _, v := range ginApps {
		v.Config()
	}

	// 完成http handler的注册
	for _, v := range ginApps {
		v.Registry(r)
	}

}

// 用户初始化 注册到Ioc容器里面的所有服务
func InitImpl() {
	for _, v := range implApps {
		v.Config()
	}
}
