package models

import (
	"boilerplate/db"
	"boilerplate/forms"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `gorm:"column:id" json:"id"`
	Email     string    `gorm:"column:email; type:varchar(255); not null; unique" json:"email"`
	Password  string    `gorm:"column:password; type:varchar(255); not null" json:"password"`
	Name      string    `gorm:"column:name; type:varchar(255); not null" json:"name"`
	RoleId    uuid.UUID `gorm:"column:role_id" json:"role_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Role      *Role     `gorm:"RoleId;references:ID" json:"role,omitempty"`
}

type UserResponseLogin struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (User) TableName() string {
	return "users"
}

type UserModel struct{}

var authModel = new(AuthModel)

func (m UserModel) Login(form forms.LoginForm) (user UserResponseLogin, token Token, err error) {
	var (
		userFetch User
	)
	err = db.GetDB().Preload("Role").Where("email = ? ", form.Email).Find(&userFetch).Error
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

func (m UserModel) Register(form forms.RegisterForm) (user User, err error) {
	getDB := db.GetDB()

	var (
		count int64
		role  Role
	)

	err = getDB.Debug().Where("email = ?", form.Email).Find(&user).Count(&count).Error
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	if count > 0 {
		return user, errors.New("email already exist")
	}

	if err = getDB.Where("name = ?", strings.ToLower(form.Role)).First(&role).Error; err != nil {
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

	err = getDB.Create(&user).Error
	if err != nil {
		return user, errors.New("something went wrong, please try again later")

	}
	return user, err
}

func (m UserModel) FindById(userID int64) (user User, err error) {
	if err = db.GetDB().First(&user, userID).Error; err != nil {
		return user, nil
	}
	return user, err
}
