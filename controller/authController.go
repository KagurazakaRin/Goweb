package controller

import (
	"github.com/gin-gonic/gin"
	"goweb/database"
	"goweb/models"
	"goweb/util"
	"net/http"
)

const localhost = "127.0.0.1:8080"
const cookieDuration = 3600 * 60 // 24h

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		panic(err)
	}

	user.SetPassword(user.Password)
	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var data, user models.User
	if err := c.ShouldBind(&data); err != nil {
		panic(err)
	}

	result := database.DB.Where("email = ?", data.Email).First(&user)

	// result.RowsAffected表示返回找到的记录数
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "the user not found",
			"info":    user.ID,
		})
		return
	}

	// 比较密码
	if err := user.ComparePassword(data.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect username or password",
		})
		return
	}

	// Set JWT and cookie
	token, err := util.GenerateJwt(user.ID, user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "JWT created failed",
		})
	}

	cookie, err := c.Cookie("jwt")
	if cookie != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "User is logged in",
		})
	} else if len(cookie) == 0 {
		c.SetCookie("jwt", token, cookieDuration, "/", localhost, false, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "User log in success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
}

func User(c *gin.Context) {
	// todo middleware

	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "not find  this user's cookie",
		})
	}

	userID, err := util.ParseJwt(cookie)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}

	var user models.User

	database.DB.Where("id = ?", userID).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"ID":       user.ID,
		"username": user.Name,
	})
}

func Logout(c *gin.Context) {
	localhost := "127.0.0.1:8080"
	c.SetCookie("jwt", "", -1, "/", localhost, false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "user is logged out",
	})
}

func Default(c *gin.Context) {
	defaultID, defaultName := 123456789, "123456789"
	token, err := util.GenerateJwt(defaultID, defaultName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "JWT created failed",
		})
	}

	cookie, err := c.Cookie("jwt")
	if cookie != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "User is logged in",
		})
	} else if len(cookie) == 0 {
		c.SetCookie("jwt", token, cookieDuration, "/", localhost, false, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "User log in success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
}
