package main

import (
	"github.com/master-of-none/go_crud/initializers"
	"github.com/master-of-none/go_crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
