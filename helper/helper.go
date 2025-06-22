package helper

import (
	"errors"

	"go_gin/models/products"
	"unicode"

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

func IsValidNameFormat(name string) bool {
	if name == "" {
		return false
	}

	first := rune(name[0])

	
	if first >= 'a' && first <= 'z' {
		return false
	}

	
	for _, r := range name {
		if r >= '0' && r <= '9' {
			return false
		}
	}

	return true
}

func IsValidPassword(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasNumber = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func VlidationProduct(p products.InsertProduct) error {

	if len(p.Name) < 3 {
		return errors.New("product name must be at least 3 characters")
	}

	if p.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	if p.Quantity <= 0 {
		return errors.New("quantity cannot be negative")
	}

	if p.Image == "" {
		return errors.New("image must be a valid URL")
	}

	if p.Category_id <= 0 {
		return errors.New("category_id is required")
	}

	return nil
}
