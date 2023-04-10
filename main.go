package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:A12345678@tcp(127.0.0.1:3306)/tddb?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return DB.Error
}

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func main() {
	// 创建数据库

	// 连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}

	// 模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", func(c *gin.Context) {
			// 添加待办事项
			var todo Todo
			c.BindJSON(&todo)
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		v1Group.GET("/todo", func(c *gin.Context) {
			// 查询待办事项
			var todolist []Todo
			if err = DB.Find(&todolist).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todolist)
			}
		})

		v1Group.GET("/todo/:id", func(c *gin.Context) {
			// 查看某个代办事项
		})

		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			// 修改某个代办事项
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效id"})
				return
			}

			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}

		})
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			// 删除某个代办事项
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效id"})
				return
			}

			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}

	// 启动
	r.Run(":9898")
}
