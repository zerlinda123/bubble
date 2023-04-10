package models

import (
	"bubble/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	Todo的增删改查操作
*/

func CreateOneTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todolist []*Todo, err error) {
	err = dao.DB.Find(&todolist).Error
	return
}

func GetOneTodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateOneTodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteOneTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(Todo{}).Error
	return err
}
