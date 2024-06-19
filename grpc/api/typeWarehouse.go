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

func (g *typeWarehouseGRPC) GetTypeInWarehouseByProductId(ctx context.Context, req *proto.GetTypeInWarehouseByProductIdReq) (*proto.GetTypeInWarehouseByProductIdRes, error) {
	var typeInWarehouses []model.TypeInWarehouse

	if err := g.db.Model(&model.TypeInWarehouse{}).Where("product_id = ?", req.ProductId).Find(&typeInWarehouses).Error; err != nil {
		return nil, err
	}

	data := []*proto.TypeInWarehouse{}

	for _, item := range typeInWarehouses {
		data = append(data, &proto.TypeInWarehouse{
			Id:          uint64(item.ID),
			ProductId:   item.ProductId,
			WarehouseId: uint64(item.WarehouseId),
			Hastag:      item.Hastag,
			Name:        item.Name,
			Price:       float32(*item.Price),
			Count:       uint64(item.Count),
			CreatedAt:   item.CreatedAt.Unix(),
			DeletedAt:   item.DeletedAt.Time.Unix(),
			UpdatedAt:   item.UpdatedAt.Unix(),
		})
	}

	res := proto.GetTypeInWarehouseByProductIdRes{
		Data: data,
	}

	return &res, nil
}

func (g *typeWarehouseGRPC) Insert(ctx context.Context, req *proto.InsertTypeInWarehouseReq) (*proto.InsertTypeInWarehouseRes, error) {
	var warehouse *model.Warehouse
	if err := g.db.Model(&model.Warehouse{}).Where("product_id = ?", req.ProductId).Find(&warehouse).Error; err != nil {
		return nil, err
	}
	var newWarehouse = model.TypeInWarehouse{
		WarehouseId: warehouse.ID,
		ProductId:   req.ProductId,
		Name:        req.Name,
		Hastag:      req.HasTag,
		Count:       uint(req.Count),
	}
	if req.Price != 0 {
		var price float64 = float64(req.Price)
		newWarehouse.Price = &price
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
