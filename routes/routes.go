package routes

import (
	"github.com/gin-gonic/gin"
	"goweb/middleware"
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
		v1Group.POST("/todo", controller.CreateTodo)

		// todo 改成 "todo/:id"
		v1Group.GET("/todo", controller.AllTodos)

		v1Group.PUT("/todo/:id", controller.UpdateTodo)

		v1Group.DELETE("/todo/:id", controller.DeleteTodo)

	}

	return r
}
