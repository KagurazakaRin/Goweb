package controller

import (
	"github.com/gin-gonic/gin"
	"goweb/database"
	"goweb/models"
	"net/http"
	"strconv"
)

func AllRoles(c *gin.Context) {
	var role []models.Role

	database.DB.Find(&role)

	c.JSON(http.StatusOK, role)
}

// CreateRole the difference between create user and register ?
// todo create 和 update 中 创建权限 permission 功能
func CreateRole(c *gin.Context) {
	var role models.Role

	if err := c.ShouldBind(&role); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}

	database.DB.Create(&role)

	c.JSON(http.StatusOK, gin.H{
		"message": role,
	})

}

func GetRole(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))
	role := models.Role{
		ID: roleID,
	}
	database.DB.Find(&role)
	c.JSON(http.StatusOK, role)
}

// UpdateRole todo error 处理 ?
func UpdateRole(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))
	role := models.Role{
		ID: roleID,
	}

	if err := c.ShouldBind(&role); err != nil {
		c.JSON(http.StatusOK, err)
	}
	database.DB.Save(&role)
	c.JSON(http.StatusOK, role)
}

func DeleteRole(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))
	role := models.Role{
		ID: roleID,
	}
	database.DB.Delete(&role)
	c.JSON(http.StatusOK, role)
}
