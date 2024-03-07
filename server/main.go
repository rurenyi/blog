package main

import (
	"blog/handle"
	"blog/handle/midware"
	"blog/utils"

	"github.com/gin-gonic/gin"

	_ "blog/docs"

	ginSwaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			博客系统
//	@version		1.0
//	@description	Swag testing
func main() {
	utils.InitLog("log")
	engine := gin.Default()
	engine.Use(midware.Cors())
	engine.GET("/swagger/*all", ginSwagger.WrapHandler(ginSwaggerFiles.Handler))
	engine.POST("/login",handle.Login)
	engine.GET("/blog/:uid",handle.GetAllBlogByName)
	engine.GET("/blog/:uid/:blogid",handle.GetBlogContent)
	engine.POST("/blog/create/:uid",handle.GreateBlog)
	engine.DELETE("/blog/delete/:uid/:blogid",handle.DeleteBlog)
	engine.Run(":12345")
}