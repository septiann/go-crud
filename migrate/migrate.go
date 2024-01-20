package main

import (
	"sept.dev/septian/go-crud/initializers"
	"sept.dev/septian/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
