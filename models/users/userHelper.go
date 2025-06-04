package users

import (
	"go_gin/helper"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModelHelper struct {
	DB *gorm.DB
}

func (u *UserModelHelper) GetAllUser() ([]Users, error) {
	users := []Users{}
	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserModelHelper) GetUser(fname, lname, email string, limit, page int) ([]Users, int64, error) {

	users := []Users{}

	var count int64

	offset := (page - 1) * limit

	if err := u.DB.Debug().Where("firstname LIKE ?", "%"+fname+"%", "lastname LIKE ?", "%"+lname+"%", "email LIKE ?", "%"+email+"%").
		Limit(limit).Offset(offset).
		Find(&users).Error; err != nil {

		return nil, 0, err
	}

	if err := u.DB.Debug().Model(&users).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (h *UserModelHelper) InsertUser(data UsersInsert) (*Users, error) {
	now := time.Now()

	user := Users{
		ID:        uuid.New().String(),
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Address:   data.Address,
		Email:     data.Email,
		Password:  helper.HashPassword(data.Password),
		CreatedAt: &now,
		UpdatedAt: &now,
		DeletedAt: nil,
	}

	if err := h.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
