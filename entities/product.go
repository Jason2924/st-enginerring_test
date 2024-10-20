package entities

type ProductSchema struct {
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" csv:"id"`
	Name          string  `gorm:"column:name;type:varchar(255)" csv:"name"`
	Price         float64 `gorm:"column:price;type:varchar(50)" csv:"price"`
	Currency      string  `gorm:"column:currency;type:varchar(10)" csv:"currency"`
	Image         string  `gorm:"column:image;type:text" csv:"image"`
	RatingAverage float64 `gorm:"column:rating_average" csv:"rating_average"`
	RatingReviews int     `gorm:"column:rating_reviews" csv:"rating_reviews"`
}

// hooks
func (ntt *ProductSchema) TableName() string {
	return "products"
}

type ProductListManyReqt struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
