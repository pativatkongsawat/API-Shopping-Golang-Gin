package products

import (
	"gorm.io/gorm"
)

type ProductModelHelper struct {
	DB *gorm.DB
}

func (u *ProductModelHelper) GetAllProducts() ([]Product, error) {
	products := []Product{}

	if err := u.DB.Debug().Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (u *ProductModelHelper) GetProduct(pname string, limit, page int) ([]Product, int64, error) {

	product := []Product{}
	var count int64
	offset := (page - 1) * limit

	if err := u.DB.Debug().Where("name LIKE ?", "%"+pname+"%").Limit(limit).Offset(offset).Find(&product).Error; err != nil {
		return nil, 0, err
	}

	if err := u.DB.Debug().Model(&product).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return product, count, nil
}
