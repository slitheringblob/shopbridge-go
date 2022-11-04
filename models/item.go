package models

type Item struct {
	Id          uint64  `json:"id" gorm:"primary-key"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price,string"`
}
