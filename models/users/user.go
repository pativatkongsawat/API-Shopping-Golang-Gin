package users

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Users struct {
	ID           string     `gorm:"primaryKey" json:"id"`
	Firstname    string     `gorm:"not null" json:"firstname"`
	Lastname     string     `gorm:"not null" json:"lastname"`
	Address      string     `json:"address"`
	Email        string     `gorm:"unique" json:"email"`
	Password     string     `gorm:"not null" json:"password"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	PermissionID int        `json:"permission_id"`
	UpdatedBy    string     `json:"updated_by"`
}

func (Users) TableName() string {
	return "users"
}

type UsersInsert struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`  
	Address   string `json:"address"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=30"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthClaims struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	Role   int    `json:"role"`
	jwt.RegisteredClaims
}

type UserUpdate struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Address   string `json:"address"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=30"`
}
