package controllers

import (
	"boilerplate/forms"
	"boilerplate/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserServiceController struct {
	Mutation models.UserMutation
}

var userForm = new(forms.UserForm)

func NewUserServiceMutation(db *gorm.DB) *UserServiceController {
	return &UserServiceController{
		Mutation: models.NewGormMutationUser(context.Background(), db),
	}
}

func (ctrl UserServiceController) Login(c *gin.Context) {
	var (
		loginForm forms.LoginForm
		ctx       = c.Request.Context()
	)

	if validationErr := c.ShouldBindJSON(&loginForm); validationErr != nil {
		message := userForm.Login(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": message,
		})
		return
	}

	user, token, err := ctrl.Mutation.Login(ctx, loginForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid Login Details",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "user": user, "token": token})
}

func (ctrl UserServiceController) Register(c *gin.Context) {
	var (
		registerForm forms.RegisterForm
		ctx          = c.Request.Context()
	)

	if err := c.ShouldBindJSON(&registerForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err,
		})
		return
	}

	user, err := ctrl.Mutation.Register(ctx, registerForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully registered",
		"Data":    user.ID,
	})
}
