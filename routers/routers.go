package routers

import (
	"github.com/gin-gonic/gin"
	"oder_query_demo/controller"
)

func SetupRouter() (r *gin.Engine) {
	r = gin.Default()
	// 定义模版文件路径
	r.LoadHTMLGlob("templates/*")
	// 定义静态文件路径
	r.Static("/static", "static")

	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		// 添加事件
		v1Group.POST("/todo", controller.AddEvent)
		// 查看所有代办事件
		v1Group.GET("/todo", controller.GetEventList)
		// 修改事件
		v1Group.PUT("/todo/:id", controller.UpdateEvent)
		// 删除事件
		v1Group.DELETE("/todo/:id", controller.DeleteEvent)
	}
	return r
}
