package middleware

import (
	"github.com/gin-gonic/gin"
	"goweb/util"
	"net/http"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		if _, err := util.ParseJwt(cookie); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "unAuthenticated",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
