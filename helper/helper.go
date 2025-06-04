package helper

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Pagination struct {
	Totalpage     int
	Prevpage      int
	Nextpage      int
	Totalrows     int
	TotalNextpage int
	Totalprevpage int
}
type LimitPage struct {
	PName string `form:"pname"`
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
}

type UserFilter struct {
	Fname string `form:"fname"`
	Lname string `form:"lname"`
	Email string `form:"email"`
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
}

func HashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashed)
}

func GenerateUUID() string {
	return uuid.New().String()
}
