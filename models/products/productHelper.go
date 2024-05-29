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

func (u *ProductModelHelper) CreateProduct([]Product) error {

	product := []Product{}
	tx := u.DB.Begin()
	if err := tx.Debug().Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (u *ProductModelHelper) UpdateProduct(products []Product) ([]Product, error) {
	tx := u.DB.Begin()

	if tx.Error != nil {
		return nil, tx.Error
	}

	updatedProducts := []Product{}

	for _, product := range products {
		updateData := map[string]interface{}{
			"Name":        product.Name,
			"Description": product.Description,
			"Price":       product.Price,
			"Quantity":    product.Quantity,
			"Image":       product.Image,
			"Update_at":   product.Update_at,
			"Category_id": product.Category_id,
		}

		if err := tx.Model(&Product{}).Where("id = ?", product.Id).Updates(&updateData).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		updatedProducts = append(updatedProducts, product)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return updatedProducts, nil
}

func (u *ProductModelHelper) DeleteProduct(id int) ([]Product, error) {

	return nil, nil
}
