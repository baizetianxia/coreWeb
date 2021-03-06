package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/baizetianxia/coreWeb/framework/gin"
	"github.com/baizetianxia/coreWeb/provider/demo"
)

func main() {
	core := gin.New()
	//绑定服务
	core.Bind(&demo.DemoServiceProvider{})

	core.Use(gin.Recovery())

	registerRouter(core)

	server := &http.Server{
		Addr:    ":8888",
		Handler: core,
	}

	// 这个goroutine是启动服务的goroutine
	go func() {
		server.ListenAndServe()
	}()

	//// 当前的 Goroutine 等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前 Goroutine 等待信号
	<-quit

	// timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// 调用Server.Shutdown graceful结束
	// if err := server.Shutdown(timeoutCtx); err != nil {
	// 	log.Fatal("Server ShutDown:", err)
	// }
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server ShutDown:", err)
	}

}
