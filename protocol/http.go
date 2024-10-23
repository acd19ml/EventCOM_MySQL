package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/acd19ml/EventCOM_MySQL/apps"
	"github.com/acd19ml/EventCOM_MySQL/conf"
	"github.com/acd19ml/EventCOM_MySQL/mcube/logger"
	"github.com/acd19ml/EventCOM_MySQL/mcube/logger/zap"
	"github.com/gin-gonic/gin"
)

// HttpService构造函数
func NewHttpService() *HttpService {
	// new gin router实例, 并没有加载Handler
	r := gin.Default()

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HttpAddr(),
		Handler:           r,
	}
	return &HttpService{
		server: server,
		l:      zap.L().Named("HTTP Service"),
		r:      r,
	}
}

type HttpService struct {
	server *http.Server
	l      logger.Logger
	r      gin.IRouter
}

func (s *HttpService) Start() error {
	// 加载handler，把所有的模块的Handler注册给了Gin Router
	apps.InitGin(s.r)

	// 已加载app的日志信息
	apps := apps.LoadedGinApps()
	s.l.Infof("loaded gin apps: %v", apps)

	// 该操作时阻塞的, 简单端口，等待请求
	// 如果服务的正常关闭, 不打印错误信息
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service stoped success")
			return nil
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}

func (s *HttpService) Stop() {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Warnf("shutdown http service error, %s", err.Error())
	}
}
