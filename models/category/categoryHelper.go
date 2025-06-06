package category

import "gorm.io/gorm"

type CategoryModelHelper struct {
	DB *gorm.DB
}

func (u *CategoryModelHelper) InsertCategory(data []Category) ([]Category, error) {

	return nil, nil
}

func (u *CategoryModelHelper) GetAllCategory() ([]Category, error) {

	categorys := []Category{}

	if err := u.DB.Find(&categorys).Error; err != nil {

		return nil, err
	}

	return categorys, nil
}
