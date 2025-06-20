package order

import "time"

type Order struct {
	Id         int        `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	CreateAt   *time.Time `json:"create_at" gorm:"create_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" gorm:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" gorm:"deleted_at"`
	UserId     string     `json:"user_id" gorm:"user_id"`
	TotalPrice float64    `json:"total_price" gorm:"total_price"`
	CreatedBy  string     `json:"created_by,omitempty" gorm:"created_by"`
	Status     string     `json:"status" gorm:"status" default:"unpaid"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderHasProduct struct {
	OrderId           int `json:"order_id" gorm:"order_id"`
	ProductId         int `json:"product_id" gorm:"product_id"`
	Id                int `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	OrderProductTotal int `json:"order_product_total" gorm:"order_product_total"`
}

func (OrderHasProduct) TableName() string {
	return "order_has_products"
}

type RequestProducts struct {
	Id       int     `json:"id"`
	Price    float64 `json:"price" gorm:"price"`
	Quantity int     `json:"quantity" gorm:"quantity"`
}

type OrderCreateRequest struct {
	Products []RequestProducts `json:"products"`
}

type Requestorder struct {
	Id         string            `json:"id" gorm:"id"`
	UserId     string            `json:"user_id"`
	TotalPrice float64           `json:"total_price" gorm:"total_price"`
	Products   []RequestProducts `json:"products"`
	Status     string            `json:"status" gorm:"status"`
}
