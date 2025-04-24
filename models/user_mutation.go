package models

import (
	"boilerplate/db"
	"boilerplate/forms"
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserMutation interface {
	Login(ctx context.Context, form forms.LoginForm) (UserResponseLogin, Token, error)
	Register(ctx context.Context, form forms.RegisterForm) (User, error)
	FindByID(ctx context.Context, userID uuid.UUID) (User, error)

	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type gormMutationUser struct {
	tx *gorm.DB
}

func NewGormMutationUser(ctx context.Context, db *gorm.DB) UserMutation {
	tx := db.WithContext(ctx).Begin()

	return &gormMutationUser{
		tx: tx,
	}
}

// Login implements UserMutation.
func (g *gormMutationUser) Login(ctx context.Context, form forms.LoginForm) (UserResponseLogin, Token, error) {
	var (
		userFetch User
		user      UserResponseLogin
		token     Token
		err       error
		authModel = new(AuthModel)
	)
	err = g.tx.Preload("Role").Where("email = ? ", form.Email).Find(&userFetch).Error
	if err != nil {
		return user, token, err
	}

	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(userFetch.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return user, token, err
	}

	//Generate the JWT auth Model
	tokenDetails, err := authModel.CreateToken(userFetch.ID, userFetch.Role.Name)
	if err != nil {
		return user, token, err
	}

	token.AccessToken = tokenDetails.AccessToken

	user.Email = userFetch.Email
	user.Name = userFetch.Name

	return user, token, nil
}

// Register implements UserMutation.
func (g *gormMutationUser) Register(ctx context.Context, form forms.RegisterForm) (User, error) {
	var (
		count int64
		role  Role
		err   error
		user  User
	)

	err = g.tx.Debug().Where("email = ?", form.Email).Find(&user).Count(&count).Error
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	if count > 0 {
		return user, errors.New("email already exist")
	}

	if err = g.tx.Where("name = ?", strings.ToLower(form.Role)).First(&role).Error; err != nil {
		return user, errors.New("role not exist")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}
	user.ID = uuid.New()
	user.Name = form.Name
	user.Email = form.Email
	user.Password = string(hashedPassword)
	user.RoleId = role.ID

	err = g.tx.Create(&user).Error
	if err != nil {
		return user, errors.New("something went wrong, please try again later")

	}
	return user, err
}

// FindByID implements UserMutation.
func (g *gormMutationUser) FindByID(ctx context.Context, userID uuid.UUID) (User, error) {
	var (
		user User
		err  error
	)
	if err = db.GetDB().First(&user, userID).Error; err != nil {
		return user, nil
	}
	return user, err
}

func (g *gormMutationUser) Commit(ctx context.Context) error {
	return g.tx.Commit().Error
}

func (g *gormMutationUser) Rollback(ctx context.Context) error {
	return g.tx.Rollback().Error
}
