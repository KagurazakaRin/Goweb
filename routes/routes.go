package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb/database"
	"goweb/middleware"
	"goweb/models"
	"net/http"

	"goweb/controller"
)

func SetRoutes() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// todo  记得给别的route添加 middleware
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// 有 cookie 的时候 ?
	r.GET("/user", middleware.IsAuthenticated(), controller.User)
	r.POST("/logout", controller.Logout)
	r.GET("/users", controller.AllUsers)
	r.GET("/users/:id", controller.GetUser)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)

	v1Group := r.Group("/v1")
	{
		// POST 新建
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo models.TodoList
			err := c.ShouldBind(&todo)

			fmt.Printf("%#v", todo)

			if err != nil {
				panic(err)
			}

			if err = database.DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    2000, // todo UPDATE magic num
					"message": "success",
					"data":    todo,
				})
			} else {
				c.JSON(http.StatusAccepted, gin.H{
					"code":    20001, // todo UPDATE magic num
					"message": "failed",
					"error":   err,
				})
			}

		})
		// GET 查看 所有/某个

		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []models.TodoList

			if err := database.DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err,
				})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})

		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})

		// UPDATE 更新状态
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			var todo models.TodoList
			id := c.Param("id")

			// todo 无效的id处理

			result := database.DB.Where("id = ?", id).First(&todo)
			if result.Error != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": result.Error,
				})
				return
			}

			err := c.ShouldBind(&todo)
			if err != nil {
				panic(err)
			}
			//DB.Save(&todo)
			database.DB.Model(&todo).Updates(map[string]interface{}{"id": todo.ID, "Title": todo.Title, "status": todo.Status})
		})
		// DELETE 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")

			// todo 无效id的处理

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
		})

	}

	return r
}
