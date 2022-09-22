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

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/logout", controller.Logout)
	r.GET("/users", controller.AllUsers)

	authed := r.Group("/")
	authed.Use(middleware.IsAuthenticated())
	{
		authed.GET("user", controller.User)
		authed.POST("users", controller.CreateUser)
		authed.GET("users/:id", controller.GetUser)
		authed.PUT("users/:id", controller.UpdateUser)
		authed.DELETE("users/:id", controller.DeleteUser)
	}

	//r.GET("/roles", controller.AllRoles)
	//r.POST("/roles", controller.CreateRole)
	//r.GET("/roles/:id", controller.GetRole)
	//r.PUT("/roles/:id", controller.UpdateRole)
	//r.DELETE("/roles/:id", controller.DeleteRole)

	r.GET("/todo", controller.AllTodos)
	todolist := r.Group("/v1")
	{
		todolist.POST("/todo", controller.CreateTodo)

		todolist.PUT("/todo/:id", controller.UpdateTodo)

		todolist.DELETE("/todo/:id", controller.DeleteTodo)

	}

	return r
}
