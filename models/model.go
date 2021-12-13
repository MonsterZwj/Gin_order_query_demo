package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID int64 `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func CreateEvent(todo *Todo) (err error) {
	err = DB.Create(&todo).Error
	return
}

func GetAllEvent() (todoList []*Todo, err error) {
	if err = DB.Find(&todoList).Error;err != nil {
		return nil, err
	}
	return
}

func GetEvent(id string) (todo *Todo, err error) {
	if err = DB.Where("id=?", id).First(&todo).Error;err != nil {
		return nil, err
	}
	return
}

func SaveEvent(todo *Todo) (err error) {
	return DB.Save(&todo).Error
}

func DeleteEvent(id string) (err error) {
	var todo Todo
	return DB.Where("id=?", id).Delete(&todo).Error
}