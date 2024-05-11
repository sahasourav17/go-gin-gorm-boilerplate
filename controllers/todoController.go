package controllers

import (
	"example/explore-go/dto"
	"example/explore-go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(ctx *gin.Context) {
	todos := models.GetTodos()
	ctx.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
}

func AddTodo(ctx *gin.Context) {
	var body dto.CreateTodoDto
	if err := ctx.BindJSON(&body); err != nil {
		log.Fatal("Error in request body parsing")
		return
	}
	result, newTodo := models.CreateTodo(body)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": newTodo,
	})
}

func GetTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := models.GetTodo(id)
	if (err) != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func UpdateTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	var body dto.UpdateTodoDto
	if err := ctx.BindJSON(&body); err != nil {
		log.Fatal("Error in request body parsing")
		return
	}
	todo, err := models.UpdateTodo(id, body)
	if (err) != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

func DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	err := models.DeleteTodo(id)
	if (err) != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
