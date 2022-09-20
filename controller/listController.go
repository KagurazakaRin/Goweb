package controller

import (
	"github.com/gin-gonic/gin"
	"goweb/database"
	"goweb/models"
	"net/http"
)

func CreateTodo(c *gin.Context) {

	var todo models.TodoList
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusOK, err)
	}

	if err := database.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    2000, // todo UPDATE magic num
			"message": "failed",
			"data":    todo,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    20001, // todo UPDATE magic num
			"message": "success",
			"error":   err,
		})
	}
}

func AllTodos(c *gin.Context) {
	var todoList []models.TodoList

	if err := database.DB.Find(&todoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context) {
	var todo models.TodoList
	id := c.Param("id")

	result := database.DB.Where("id = ?", id).First(&todo)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": result.Error,
		})
		return
	}

	err := c.ShouldBind(&todo)

	c.JSON(http.StatusOK, gin.H{
		"error": err,
	})

	database.DB.Model(&todo).Updates(map[string]interface{}{"id": todo.ID, "Title": todo.Title, "status": todo.Status})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	result := database.DB.Where("id = ?", id).Delete(&models.TodoList{})
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": "Delete",
		})
	}
}
