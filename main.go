package main

import (
	"oder_query_demo/models"
	"oder_query_demo/routers"
)

func main() {
	// 连接数据库
	err := models.InitMySQL()
	if err != nil {
		panic(err)
	}
	err = models.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		return
	}
	r := routers.SetupRouter()
	r.Run(":8080")
}
