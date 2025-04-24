package controllers

import (
	"boilerplate/domain"
	"boilerplate/forms"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserServiceController struct {
	DB *gorm.DB
}

var userForm = new(forms.UserForm)

func NewUserServiceMutation(db *gorm.DB) *UserServiceController {
	return &UserServiceController{
		DB: db,
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
	mutation := domain.NewGormMutationUser(ctx, ctrl.DB)
	user, token, err := mutation.Login(ctx, loginForm)
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
	mutation := domain.NewGormMutationUser(ctx, ctrl.DB)

	user, err := mutation.Register(ctx, registerForm)
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully registered",
		"Data":    user.ID,
	})
}
