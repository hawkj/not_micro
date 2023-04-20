package main

import (
	"context"
	"gitee.com/lichuan2022/my-todo/config"
	"gitee.com/lichuan2022/my-todo/pkg/common"
	"gitee.com/lichuan2022/my-todo/pkg/db_mysql/my_todo"
	"gitee.com/lichuan2022/my-todo/pkg/handler/mytodo"
	"gitee.com/lichuan2022/my-todo/pkg/middleware"
	"gitee.com/lichuan2022/my-todo/pkg/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 创建Gin引擎
	r := gin.Default()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	//export APP_ENV=dev
	env := os.Getenv("APP_ENV")
	configFile := ""
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
		configFile = "/config/conf.yaml"
	} else {
		gin.SetMode(gin.DebugMode)
		configFile = "/config/test_conf.yaml"
	}
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config := config.GetConfig(currentDir + configFile)

	g := &common.Global{
		DbMyTodo: mytododb.NewDb(config.MyTodoDb),
		Redis:    redis.NewRedis(config.Redis),
	}

	baseGroup := r.Group("", middleware.CommonContext(g, config))

	commonApi := baseGroup.Group("/api")
	commonApi.GET("/test", middleware.MyHandlerWrapper(handlermytodo.Test))

	taskApi := baseGroup.Group("/api/task")
	taskApi.Use(middleware.TimeoutMiddleware(3 * time.Second))

	taskApi.POST("/create", middleware.MyHandlerWrapper(handlermytodo.TaskCreate))
	taskApi.POST("/msg_post", middleware.MyHandlerWrapper(handlermytodo.PostMsg))
	// 启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待停止信号
	<-ctx.Done()

	// 设置超时时间
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭 HTTP 服务器
	log.Println("shutting down server...")
	if err := srv.Shutdown(timeoutCtx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("server exiting")
}
