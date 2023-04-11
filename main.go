package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {
	// 创建数据库
	// sql: Create DATABASE tddb;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	// 注册路由
	r := routers.SetupRouter()

	// 启动
	r.Run(":9898")
}
