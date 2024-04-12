package model

import "gorm.io/gorm"

type TypeInWarehouse struct {
	gorm.Model
	ProductId   string  `json:"productId"`
	WarehouseId uint    `json:"warehouseId" gorm:"uniqueIndex:idx_wh_ht"`
	Hastag      string  `json:"hastag" gorm:"uniqueIndex:idx_wh_ht"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Count       uint    `json:"count"`
}
