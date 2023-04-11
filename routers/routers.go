package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.GET("/index", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.AddTodo)

		v1Group.GET("/todo", controller.GetTodo)

		v1Group.GET("/todo/:id", func(c *gin.Context) {
			// 查看某个代办事项
		})

		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
