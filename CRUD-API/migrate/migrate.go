package main

import (
	"github.com/ManavMuthanna/CRUD-API/initializers"
	"github.com/ManavMuthanna/CRUD-API/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
