package controller

import (
	"github.com/gin-gonic/gin"
	"goweb/database"
	"goweb/models"
	"net/http"
	"strconv"
)

func AllUsers(c *gin.Context) {
	var user []models.User

	database.DB.Find(&user)

	c.JSON(http.StatusOK, user)
}

// the difference between create user and register ?
// func CreateUser(c *gin.Context) {}

func GetUser(c *gin.Context) {
	//fmt.Println("success")
	userID, _ := strconv.Atoi(c.Param("id"))
	user := models.User{
		ID: userID,
	}
	database.DB.Find(&user)
	c.JSON(http.StatusOK, user)
}

// UpdateUser todo error 处理 ?
func UpdateUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	user := models.User{
		ID: userID,
	}

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, err)
	}
	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	user := models.User{
		ID: userID,
	}
	database.DB.Delete(&user)
	c.JSON(http.StatusOK, user)
}
