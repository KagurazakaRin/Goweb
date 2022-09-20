package main

import (
	"goweb/database"
	"goweb/routes"
)

// todo init mysql

func main() {

	database.Connect()
	/*
		dsn := "root:windows@(127.0.0.1:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Could not connect to the database")
		}
	*/

	r := routes.SetRoutes()
	/*
		r := gin.Default()

		r.Static("/static", "static")
		r.LoadHTMLGlob("templates/*")

		r.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})

		v1Group := r.Group("/v1")
		{
			// POST 新建
			v1Group.POST("/todo", func(c *gin.Context) {
				var todo TodoList
				err = c.ShouldBind(&todo)

				fmt.Printf("%#v", todo)

				if err != nil {
					panic(err)
				}

				if err = DB.Create(&todo).Error; err != nil {
					c.JSON(http.StatusOK, gin.H{
						"code":    DatabaseCreated,
						"message": "success",
						"data":    todo,
					})
				} else {
					c.JSON(http.StatusAccepted, gin.H{
						"code":    DatabaseNotCreated,
						"message": "failed",
						"error":   err,
					})
				}

			})
			// GET 查看 所有/某个

			v1Group.GET("/todo", func(c *gin.Context) {
				var todoList []TodoList

				if err = DB.Find(&todoList).Error; err != nil {
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
				var todo TodoList
				id := c.Param("id")

				// todo 无效的id处理

				result := DB.Where("id = ?", id).First(&todo)
				if result.Error != nil {
					c.JSON(http.StatusOK, gin.H{
						"error": result.Error,
					})
					return
				}

				err = c.ShouldBind(&todo)
				if err != nil {
					panic(err)
				}
				//DB.Save(&todo)
				DB.Model(&todo).Updates(map[string]interface{}{"id": todo.ID, "Title": todo.Title, "status": todo.Status})
			})
			// DELETE 删除
			v1Group.DELETE("/todo/:id", func(c *gin.Context) {
				id := c.Param("id")

				// todo 无效id的处理

				result := DB.Where("id = ?", id).Delete(&TodoList{})
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
	*/
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
