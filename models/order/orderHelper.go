package order

import (
	"gorm.io/gorm"
)

type OrderModelHelper struct {
	DB *gorm.DB
}

func (u *OrderModelHelper) CreateOrder(data *Order) (*Order, error) {

	tx := u.DB.Begin()
	if err := tx.Debug().Create(&data).Error; err != nil {

		tx.Rollback()
		return nil, err

	}
	tx.Commit()

	return data, nil
}

func (u *OrderModelHelper) CreateOrderHasProduct(orderId int, data []RequestProducts) (*[]OrderHasProduct, error) {

	orderdetail := []OrderHasProduct{}
	tx := u.DB.Begin()

	for _, p := range data {

		newdata := OrderHasProduct{
			OrderId:           orderId,
			ProductId:         p.Id,
			OrderProductTotal: p.Quantity,
		}

		if err := tx.Debug().Table("order_has_products").Create(&newdata).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		orderdetail = append(orderdetail, newdata)
	}

	return &orderdetail, nil
}
