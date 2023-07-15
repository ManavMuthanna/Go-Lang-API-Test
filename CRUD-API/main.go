package main

import (
	"github.com/ManavMuthanna/CRUD-API/controllers"
	"github.com/ManavMuthanna/CRUD-API/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/post/:id", controllers.PostsShow)
	r.PUT("/post/:id", controllers.PostsUpdate)
	r.DELETE("/post/:id", controllers.PostsDelete)

	r.Run("localhost:3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
