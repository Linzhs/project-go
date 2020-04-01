package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	log2 "project-go/pkg/log"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"

	_ "go.uber.org/automaxprocs"
)

const (
	serverAddr = ":8888"
)

func main() {
	logger, err := zap.NewDevelopment(zap.AddCaller())
	if err != nil {
		panic(err)
	}

	appEngine := gin.New()
	appEngine.Use(log2.RegisterGinZap(logger, time.RFC3339, false)).Use(gin.Recovery())
	appEngine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	appEngine.Run(":8080")

	srv := &http.Server{
		Addr:    serverAddr,
		Handler: appEngine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待系统的中断信号来gracefully shutdown服务, 这里设置30秒来下线服务
	quit := make(chan os.Signal)
	// kill 没有指定 '-' 参数时 默认发送syscall.SIGTERM
	// kill -2 为 syscall.SIGINT
	// kill -9 为 syscall.SIGKILL 这个信号不能被捕捉 所以不需要添加这个
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server now ...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown: ", err)
	}

	select {
	case <-ctx.Done():
		log.Println("shutdown context timeout.")
	}

	log.Println("Server exiting")
}
