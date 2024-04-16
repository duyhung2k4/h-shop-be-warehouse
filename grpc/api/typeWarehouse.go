package api

import (
	"app/config"
	"app/grpc/proto"
	"app/model"
	"context"

	"gorm.io/gorm"
)

type typeWarehouseGRPC struct {
	db *gorm.DB
	proto.UnsafeTypeInWarehouseServiceServer
}

func (g *typeWarehouseGRPC) Insert(ctx context.Context, req *proto.InsertTypeInWarehouseReq) (*proto.InsertTypeInWarehouseRes, error) {
	var newWarehouse = model.TypeInWarehouse{
		ProductId: req.ProductId,
	}

	if err := g.db.Model(&model.TypeInWarehouse{}).Create(&newWarehouse).Error; err != nil {
		return nil, err
	}

	res := &proto.InsertTypeInWarehouseRes{
		Id:        uint64(newWarehouse.ID),
		ProductId: newWarehouse.ProductId,
		Count:     uint64(newWarehouse.Count),
	}

	return res, nil
}

func (g *typeWarehouseGRPC) Update(ctx context.Context, req *proto.UpdateTypeInWarehouseReq) (*proto.UpdateTypeInWarehouseRes, error) {
	newWarehouse := model.TypeInWarehouse{
		Count: uint(req.Count),
	}

	if err := g.db.
		Model(&model.TypeInWarehouse{}).
		Where("id = ? AND product_id = ?", req.Id).
		Updates(&newWarehouse).Error; err != nil {
		return nil, err
	}

	res := &proto.UpdateTypeInWarehouseRes{
		Id:        uint64(newWarehouse.ID),
		ProductId: newWarehouse.ProductId,
		Count:     uint64(newWarehouse.Count),
	}

	return res, nil
}

func (g *typeWarehouseGRPC) UpCount(ctx context.Context, req *proto.UpCountTypeInWarehouseReq) (*proto.UpCountTypeInWarehouseRes, error) {
	if err := g.db.
		Model(&model.TypeInWarehouse{}).
		Where("id = ?", req.Id).
		UpdateColumn("count", gorm.Expr("count + ?", req.Amount)).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (g *typeWarehouseGRPC) DownCount(ctx context.Context, req *proto.DownCountTypeInWarehouseReq) (*proto.DownCountTypeInWarehouseRes, error) {
	if err := g.db.
		Model(&model.TypeInWarehouse{}).
		Where("id = ?", req.Id).
		UpdateColumn("count", gorm.Expr("count - ?", req.Amount)).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func NewTypeInWarehouseGRPC() proto.TypeInWarehouseServiceServer {
	return &typeWarehouseGRPC{
		db: config.GetDB(),
	}
}