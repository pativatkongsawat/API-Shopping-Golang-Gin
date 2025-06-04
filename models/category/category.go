package category

type Category struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name string `json:"name" gorm:"name"`
}

func (Category) TableName() string {

	return "categories"

}
