package category

import (
	"gorm.io/gorm"
)

type CategoryModelHelper struct {
	DB *gorm.DB
}

func (u *CategoryModelHelper) InsertCategory(data []Category) ([]Category, error) {
	var categorydata []Category

	for _, d := range data {
		category := Category{
			Name: d.Name,
		}

		categorydata = append(categorydata, category)
	}

	if err := u.DB.Create(&categorydata).Error; err != nil {
		return nil, err
	}

	return categorydata, nil
}

func (u *CategoryModelHelper) GetAllCategory() ([]Category, error) {

	categorys := []Category{}

	tx := u.DB.Begin()

	if err := tx.Find(&categorys).Error; err != nil {

		return nil, err
	}

	return categorys, nil
}
