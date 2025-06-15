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

func (u *CategoryModelHelper) DeleteCategory(id int) ([]Category, error) {

	category := []Category{}

	tx := u.DB.Begin()

	if err := tx.Debug().Where("id ? = ", id).Delete(&category).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return category, nil
}

func (u *CategoryModelHelper) UpdateCategory(data []Category) ([]Category, error) {

	updateCategory := []Category{}
	tx := u.DB.Begin()

	for _, category := range data {
		updatadata := map[string]interface{}{
			"Name": category.Name,
		}
		if err := tx.Debug().Model(&Category{}).Where("id = ?", category.Id).Updates(updatadata).Error; err != nil {
			tx.Rollback()
			return nil, err

		}

		updateCategory = append(updateCategory, category)
	}
	tx.Commit()

	return updateCategory, nil
}
