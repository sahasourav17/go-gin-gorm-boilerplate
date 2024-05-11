package models

import (
	"errors"
	"example/explore-go/config"
	"example/explore-go/dto"
	"example/explore-go/schemas"
	"fmt"

	"gorm.io/gorm"
)

func CreateTodo(fields dto.CreateTodoDto) (*gorm.DB, schemas.Todo) {
	newTodo := schemas.Todo{Title: fields.Title}
	result := config.DB.Create(&newTodo)
	return result, newTodo
}

func GetTodos() []schemas.Todo {
	var todos []schemas.Todo
	config.DB.Find(&todos)

	return todos
}

func GetTodo(todoId string) (schemas.Todo, error) {
	var todo schemas.Todo
	if err := config.DB.First(&todo, todoId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return todo, fmt.Errorf("todo with ID %s not found", todoId)
		}
		return todo, err
	}
	return todo, nil
}

func UpdateTodo(todoId string, fields dto.UpdateTodoDto) (schemas.Todo, error) {
	todo, fetchErr := GetTodo(todoId)
	if fetchErr != nil {
		return todo, fetchErr
	}

	config.DB.Model(&todo).Updates(fields)
	return todo, nil
}

func DeleteTodo(todoId string) error {
	todo, fetchErr := GetTodo(todoId)
	if fetchErr != nil {
		return fetchErr
	}

	config.DB.Delete(&todo)
	return nil
}
