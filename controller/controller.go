package controller

import (
	"bubble/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func AddTodo(c *gin.Context) {
	// 添加待办事项
	var todo models.Todo
	c.BindJSON(&todo)
	if err := models.CreateOneTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodo(c *gin.Context) {
	// 查询待办事项

	todolist, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func UpdateTodo(c *gin.Context) {
	// 修改某个代办事项
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}

	todo, err := models.GetOneTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	c.BindJSON(&todo)
	if err := models.UpdateOneTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func DeleteTodo(c *gin.Context) {
	// 删除某个代办事项
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}

	if err := models.DeleteOneTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
