package users

import (
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
