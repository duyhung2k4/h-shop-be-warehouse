package model

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	ProductId string `json:"productId"`
	Count     uint   `json:"count" gorm:"default:0"`

	TypeInWarehouses []TypeInWarehouse `json:"typeInWarehouses" gorm:"foreignKey:WarehouseId"`
}
