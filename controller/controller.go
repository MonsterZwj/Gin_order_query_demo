package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oder_query_demo/models"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func AddEvent(c *gin.Context) {
	// 获取前端传过来的参数
	var todo models.Todo
	c.BindJSON(&todo)
	// 写入数据库
	err := models.CreateEvent(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"data": todo,
		})
	}
}

func GetEventList(c *gin.Context) {
	if todoList, err := models.GetAllEvent();err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateEvent(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	todo, err := models.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&todo)
	if err = models.SaveEvent(todo);err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteEvent(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	if err := models.DeleteEvent(id);err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{id: "success"})
	}
}
