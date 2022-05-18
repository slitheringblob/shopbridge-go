package models

type Item struct {
	id          uint64  `json:"id" gorm:"primary-key"`
	name        string  `json:"name"`
	description string  `json:"description"`
	price       float64 `json:"price"`
}
