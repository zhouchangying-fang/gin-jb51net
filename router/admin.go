package router

import (
	"gin/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.LoadHTMLGlob("templates/admin/**/*")
	r.Static("/static", "./public/static")

	admin := r.Group("/admin")
	{
		admin.Use(controller.Base{}.BaseController)
		admin.GET("/", controller.Admin{}.Index)
		admin.GET("/user", controller.Admin{}.User)
	}
	r.GET("/admin/login", controller.Login{}.Index)
	r.POST("/admin/login/load", controller.Login{}.Index)
	return r
}
