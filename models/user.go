package models

import (
	"boilerplate/forms"
	"time"

	"github.com/google/uuid"
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

func (u *User) CreateNewUser(form forms.RegisterForm, hashedPassword []byte, roleID uuid.UUID) {
	u.ID = uuid.New()
	u.Name = form.Name
	u.Email = form.Email
	u.Password = string(hashedPassword)
	u.RoleId = roleID
}
