package main

import (
	config "MyBlog"
	"MyBlog/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	//config.DB.AutoMigrate(&model.User{})
	r := gin.Default()

	//用户注册路由
	r.POST("/api/v1/users", controllers.Register)
	r.POST("/api/v1/users/login", controllers.Login)
	//r.GET("/api/v1/users/:id", controllers.GetUser)
	//r.PUT("/api/v1/users/:id", controllers.UpdateUser)
	r.Run(":8080")
}
