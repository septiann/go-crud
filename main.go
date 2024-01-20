package main

import (
	"github.com/gin-gonic/gin"
	"sept.dev/septian/go-crud/controllers"
	"sept.dev/septian/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/post", controllers.PostCreate)
	r.GET("/post", controllers.PostList)
	r.GET("/post/:id", controllers.PostShow)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()
}
