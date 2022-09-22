package controller

import (
	"github.com/gin-gonic/gin"
	"goweb/database"
	"goweb/models"
	"math"
	"net/http"
	"strconv"
)

func AllUsers(c *gin.Context) {
	var user []models.User
	var total int64

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	database.DB.Preload("Role").Offset(offset).Limit(limit).Find(&user)
	database.DB.Model(&models.User{}).Count(&total)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total) / float64(limit)),
		},
	})
}

// CreateUser the difference between create user and register ?
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})

}

func GetUser(c *gin.Context) {
	//fmt.Println("success")
	userID, _ := strconv.Atoi(c.Param("id"))
	user := models.User{
		ID: userID,
	}
	database.DB.Preload("Role").Find(&user)
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
