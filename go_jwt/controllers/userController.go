package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/master-of-none/go_jwt/initializers"
	"github.com/master-of-none/go_jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	// Get the EMAIL and password of req body
	var body struct {
		Email    string
		Password string
	}

	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read Body",
		})
		return
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash the password",
		})
		return
	}
	// Create the user

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to creat the User",
		})
		return
	}
	// Send a response
	ctx.JSON(http.StatusOK, gin.H{})
}
