package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"bubble/setting"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main config/config.ini")
		return
	}

	err := setting.InitConf(os.Args[1])
	if err != nil {
		fmt.Printf("read config.ini failed,err=%v\n", err)
		return

	}
	// 创建数据库
	// sql: Create DATABASE tddb;
	// 连接数据库
	err = dao.InitMySQL(setting.Conf.MySQLConf)
	if err != nil {
		fmt.Printf("init MySQL failed,err=%v\n", err)
		return
	}

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	// 注册路由
	r := routers.SetupRouter()

	// 启动
	err = r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("server start failed,err=%v\n", err)
	}
}
