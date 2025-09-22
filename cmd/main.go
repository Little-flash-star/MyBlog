package main

import (
	"MyBlog/config"
	"MyBlog/internal/controllers"
	"MyBlog/internal/middleware"
	"MyBlog/internal/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	config.DB.AutoMigrate(&models.User{})
	r := gin.Default()
	r.Use(cors.Default())
	//用路由
	r.POST("/api/v1/users", controllers.Register)
	r.POST("/api/v1/users/login", controllers.Login)
	//r.GET("/api/v1/users/:id", controllers.GetUser)
	//r.PUT("/api/v1/users/:id", controllers.UpdateUser)

	articleGroup := r.Group("/articles")
	{
		// 创建、更新、删除需要登录
		articleGroup.POST("", middleware.JWTAuthMiddleware(), controllers.CreateArticle)
		articleGroup.PUT("/:id", middleware.JWTAuthMiddleware(), controllers.UpdateArticle)
		articleGroup.DELETE("/:id", middleware.JWTAuthMiddleware(), controllers.DeleteArticle)

		// 查询文章列表和详情无需登录
		articleGroup.GET("", controllers.GetArticles)
		articleGroup.GET("/:id", controllers.GetArticle)
	}

	r.Run(":8080")
}
