package models

import "time"

type Fridge struct {
	ID int `json:"id" gorm:"primary_key; auto_increment:true"`
	Name string `gorm:"size:100; UNIQUE_INDEX"`
	Groceries []Grocery
}

type Grocery struct {
	ID int `gorm:"primary_key; auto_increment:true" json:"id"`
	Name string `gorm:"size:100; UNIQUE_INDEX" json:"name"`
	Quantity float64 `json:"quantity"`
	Unit string `gorm:"size:50" json:"unit"`
	Category string `gorm:"size:50" json:"category"`
	Expiry time.Time `json:"expiry" gorm:"type: date"`
	FridgeID int `json:"fridgeId"`
}
