package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"goweb/database"
	"goweb/models"
	"goweb/util"
	"net/http"
)

const passwordCost = 14

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		panic(err)
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), passwordCost)
	user.Password = string(hashPassword)

	database.DB.Create(&user)

	c.JSON(http.StatusOK, user)

}

func Login(c *gin.Context) {
	var data, user models.User
	if err := c.ShouldBind(&data); err != nil {
		panic(err)
	}

	result := database.DB.Where("email = ?", data.Email).First(&user)
	// 如果在数据库中无法找到该email，（email唯一）
	// result.RowsAffected表示返回找到的记录数
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "the user not found",
			"info":    user,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect username or password",
		})
		return
	}

	// util.GenerateJwt(id, name) id : 1, name : "jwt"
	token, err := util.GenerateJwt(1)

	/*
		// todo claim 改一下里面的内容；改完以后记得修改 User里面的jwt.Parse
		claims := &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    strconv.Itoa(user.ID),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		signedToken, err := token.SignedString([]byte(SigningKey))
	*/

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
			"msg":   "JWT created failed",
		})
	}

	_, err = c.Cookie("jwt")
	if err != nil {
		//cookie = "NotSet"
		localhost := "127.0.0.1:8080"
		c.SetCookie("jwt", token, 3600*60, "/", localhost, false, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}

	// SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool)
	// MaxAge设置为-1，表示删除cookie; 默认好像是GMT时间，伦敦
	c.JSON(http.StatusOK, user)
	c.JSON(http.StatusOK, gin.H{
		"signedToken": token,
	})

}

func User(c *gin.Context) {
	//fmt.Println("get user")

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
		"claim": user,
	})
}

func Logout(c *gin.Context) {
	localhost := "127.0.0.1:8080"
	c.SetCookie("jwt", "", -1, "/", localhost, false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}