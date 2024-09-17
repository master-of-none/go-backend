package main

import (
	"github.com/gin-gonic/gin"
	"github.com/master-of-none/go_jwt/controllers"
	"github.com/master-of-none/go_jwt/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello There",
		})
	})
	r.POST("/signup", controllers.SignUp)

	r.Run()
}
