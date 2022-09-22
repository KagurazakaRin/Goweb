package controller

import (
	"github.com/gin-gonic/gin"
	"goweb/database"
	"goweb/models"
	"goweb/util"
	"net/http"
	"strconv"
)

func CreateTodo(c *gin.Context) {

	var todo models.TodoList
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusOK, err)
	}

	cookie, _ := c.Cookie("jwt")
	userID, _ := util.ParseJwt(cookie)
	todo.Uid, _ = strconv.Atoi(userID)

	if err := database.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "failed to create todo",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Uid":     todo.Uid,
			"id":      todo.ID,
			"message": "Successfully created",
		})
	}
}

func AllTodos(c *gin.Context) {

	cookie, _ := c.Cookie("jwt")
	userID, _ := util.ParseJwt(cookie)
	uid, _ := strconv.Atoi(userID)

	var todoList []models.TodoList

	if err := database.DB.Preload("User").Where("uid = ?", uid).Find(&todoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error message": "doesn't have permission",
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context) {
	var todo models.TodoList
	id := c.Param("id")

	cookie, _ := c.Cookie("jwt")
	userID, _ := util.ParseJwt(cookie)
	uid, _ := strconv.Atoi(userID)

	result := database.DB.Preload("User").Where("id = ? AND uid = ?", id, uid).First(&todo)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error message": "Not found this todo id or not have permission",
		})
		return
	}

	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error message": "bind failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Updated todo",
		})
	}

	database.DB.Model(&todo).Updates(map[string]interface{}{"id": todo.ID, "Title": todo.Title, "status": todo.Status})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	cookie, _ := c.Cookie("jwt")
	userID, _ := util.ParseJwt(cookie)
	uid, _ := strconv.Atoi(userID)

	result := database.DB.Where("id = ? AND uid = ?", id, uid).Delete(&models.TodoList{})
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error message": "Not found this todo id or not have permission",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": "The todo deleted",
		})
	}
}
