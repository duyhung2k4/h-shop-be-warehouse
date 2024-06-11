package model

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	ProductId string `json:"productId" gorm:"uniqueIndex:idx_product_id"`
	Count     uint   `json:"count" gorm:"default:0;check:check_count, count >= 0"`

	TypeInWarehouses []TypeInWarehouse `json:"typeInWarehouses" gorm:"foreignKey:WarehouseId"`
}
