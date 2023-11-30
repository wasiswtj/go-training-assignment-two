package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"orderId"`

	Order *Order `json:"-"`
}
