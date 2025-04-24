package controllers

import (
	"boilerplate/forms"
	"boilerplate/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.UserModel)
var userForm = new(forms.UserForm)

func (ctrl UserController) Login(c *gin.Context) {
	var loginForm forms.LoginForm

	if validationErr := c.ShouldBindJSON(&loginForm); validationErr != nil {
		message := userForm.Login(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": message,
		})
		return
	}

	user, token, err := userModel.Login(loginForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid Login Details",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "user": user, "token": token})
}

func (ctrl UserController) Register(c *gin.Context) {
	var registerForm forms.RegisterForm

	if err := c.ShouldBindJSON(&registerForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err,
		})
		return
	}
	log.Println(registerForm)
	user, err := userModel.Register(registerForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully registered",
		"user":    user,
	})
}
