package main

import (
	"github.com/RODO94/go-crud-app/initializers"
	"github.com/RODO94/go-crud-app/models"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	initializers.DB.AutoMigrate(&models.Post{})
}