package users

import (
	"errors"
	"fmt"
	"go_gin/helper"

	"time"

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

func (u *UserModelHelper) InsertUser(data []UsersInsert) ([]Users, error) {
	var usersToInsert []Users
	now := time.Now()

	for _, d := range data {
		if !helper.IsValidPassword(d.Password) {
			return nil, fmt.Errorf("password must be at least 8 characters and include uppercase, lowercase, number, and special character")
		}

		var existing Users
		if err := u.DB.Where("email = ?", d.Email).First(&existing).Error; err == nil {
			return nil, fmt.Errorf("email %s already exists", d.Email)
		}

		user := Users{
			ID:        helper.GenerateUUID(),
			Firstname: d.Firstname,
			Lastname:  d.Lastname,
			Address:   d.Address,
			Email:     d.Email,
			Password:  helper.HashPassword(d.Password),
			CreatedAt: &now,
			UpdatedAt: &now,
			DeletedAt: nil,
			PermissionID: 2,
		}
		usersToInsert = append(usersToInsert, user)
	}

	if err := u.DB.Create(&usersToInsert).Error; err != nil {
		return nil, err
	}

	return usersToInsert, nil
}

func (u *UserModelHelper) Register(data []Users) ([]Users, error) {

	tx := u.DB.Begin()
	var emails []string
	for _, user := range data {
		emails = append(emails, user.Email)

		if !helper.IsValidPassword(user.Password) {

			tx.Rollback()
			return nil, fmt.Errorf("password must be at least 8 characters and include uppercase, lowercase, number, and special character")
		}

		if !helper.IsValidNameFormat(user.Firstname) {
			tx.Rollback()
			return nil, fmt.Errorf("invalid first name format")
		}
		if !helper.IsValidNameFormat(user.Lastname) {
			tx.Rollback()
			return nil, fmt.Errorf("invalid last name format")
		}

	}

	var existingUsers []Users
	if err := u.DB.Where("email IN ?", emails).Find(&existingUsers).Error; err != nil {
		return nil, err
	}

	if len(existingUsers) > 0 {
		return nil, errors.New("some emails already exist")
	}

	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return data, nil
}

func (u *UserModelHelper) UpdateUser(useremail string, data []UserUpdate) ([]Users, error) {
	now := time.Now()
	updatedUsers := []Users{}
	tx := u.DB.Begin()

	for _, p := range data {
		if !helper.IsValidPassword(p.Password) {
			tx.Rollback()
			return nil, fmt.Errorf("password must be at least 8 characters and include uppercase, lowercase, number, and special character")
		}

		if !helper.IsValidNameFormat(p.Firstname) {
			tx.Rollback()
			return nil, fmt.Errorf("invalid first name format")
		}

		if !helper.IsValidNameFormat(p.Lastname) {
			tx.Rollback()
			return nil, fmt.Errorf("invalid last name format")
		}

		updateData := map[string]interface{}{
			"Firstname": p.Firstname,
			"Lastname":  p.Lastname,
			"Address":   p.Address,
			"Email":     p.Email,
			"Password":  helper.HashPassword(p.Password),
			"UpdatedAt": &now,
			"UpdatedBy": useremail,
		}

		if err := tx.Model(&Users{}).Where("email = ?", p.Email).Updates(updateData).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		var updatedUser Users
		if err := tx.Where("email = ?", p.Email).First(&updatedUser).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		updatedUsers = append(updatedUsers, updatedUser)
	}

	tx.Commit()
	return updatedUsers, nil
}
