package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"wps_go/service_discovery"
	"wps_go/user_service/controllers"
	"wps_go/user_service/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	database.InitDB()

	r := gin.Default()

	// 用户服务路由
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
		userGroup.GET("/info", controllers.GetUserInfo)
		userGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}

	// 注册服务到Consul
	err := service_discovery.RegisterService("user-service", 8080)
	if err != nil {
		log.Fatal("Failed to register service: ", err)
	}
	defer func() {
		if err := service_discovery.DeregisterService("user-service"); err != nil {
			log.Println("Failed to deregister service: ", err)
		}
	}()

	// 优雅退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")
	}()

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
