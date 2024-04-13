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

func (g *warehouseGRPC) Insert(ctx context.Context, req *proto.InsertWarehouseReq) (*proto.InsertWarehouseRes, error) {
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

func (g *warehouseGRPC) Update(ctx context.Context, req *proto.UpdateWarehouseReq) (*proto.UpdateWarehouseRes, error) {
	newWarehouse := model.Warehouse{
		Count: uint(req.Count),
	}

	if err := g.db.
		Model(&model.Warehouse{}).
		Where("id = ? AND product_id = ?", req.Id).
		Updates(&newWarehouse).Error; err != nil {
		return nil, err
	}

	res := &proto.UpdateWarehouseRes{
		Id:        uint64(newWarehouse.ID),
		ProductId: newWarehouse.ProductId,
		Count:     uint64(newWarehouse.Count),
	}

	return res, nil
}

func (g *warehouseGRPC) UpCount(ctx context.Context, req *proto.UpCountWarehouseReq) (*proto.UpCountWarehouseRes, error) {
	if err := g.db.
		Model(&model.Warehouse{}).
		Where("id = ?", req.Id).
		UpdateColumn("count", gorm.Expr("count + ?", 1)).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (g *warehouseGRPC) DownCount(ctx context.Context, req *proto.DownCountWarehouseReq) (*proto.DownCountWarehouseRes, error) {
	if err := g.db.
		Model(&model.Warehouse{}).
		Where("id = ?", req.Id).
		UpdateColumn("count", gorm.Expr("count - ?", 1)).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func NewWarehouseGRPC() proto.WarehouseServiceServer {
	return &warehouseGRPC{
		db: config.GetDB(),
	}
}
