package api

import (
	"app/config"
	"app/grpc/proto"
	"app/model"
	"context"

	"gorm.io/gorm"
)

type warehouseGRPC struct {
	db *gorm.DB
	proto.UnsafeWarehouseServiceServer
}

func (g *warehouseGRPC) InsertWarehouse(ctx context.Context, req *proto.InsertWarehouseReq) (*proto.InsertWarehouseRes, error) {
	var newWarehouse = model.Warehouse{
		ProductId: req.ProductId,
	}

	if err := g.db.Model(&model.Warehouse{}).Create(&newWarehouse).Error; err != nil {
		return nil, err
	}

	res := &proto.InsertWarehouseRes{
		Id:        uint64(newWarehouse.ID),
		ProductId: newWarehouse.ProductId,
		Count:     uint64(newWarehouse.Count),
	}

	return res, nil
}

func NewWarehouseGRPC() proto.WarehouseServiceServer {
	return &warehouseGRPC{
		db: config.GetDB(),
	}
}
