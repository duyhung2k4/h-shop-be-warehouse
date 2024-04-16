package api

import (
	"app/config"
	"app/grpc/proto"
	"app/model"
	"context"
	"sync"

	"gorm.io/gorm"
)

type countPriceGRPC struct {
	db *gorm.DB
	proto.UnsafeCountPriceServiceServer
}

func (g *countPriceGRPC) CountPriceOrder(ctx context.Context, req *proto.CountPriceOrderReq) (*proto.CountPriceOrderRes, error) {
	var errGetData error = nil
	var price float64

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, item := range req.Orders {
		wg.Add(1)
		go func(order *proto.Order) {
			if item.TypeInWarehouseId == 0 {
				var warehouse *model.Warehouse
				err := g.db.
					Model(&model.Warehouse{}).
					Where("id = ?", item.WarehouseId).
					First(&warehouse).
					Error

				if err != nil {
					errGetData = err
				}

				if warehouse != nil {

				}
			} else {
				var typeInWarehouse *model.TypeInWarehouse
				err := g.db.
					Model(&model.TypeInWarehouse{}).
					Preload("Warehouse").
					Where("id = ?", item.TypeInWarehouseId).
					First(&typeInWarehouse).
					Error

				if err != nil {
					errGetData = err
				}

				if typeInWarehouse != nil {
					mutex.Lock()
					if typeInWarehouse.Price != nil {
						price += *typeInWarehouse.Price * float64(order.Amount)
					} else {

					}
					mutex.Unlock()
				}
			}

			wg.Done()
		}(item)
	}

	wg.Wait()

	return nil, nil
}

func NewCountPriceGRPC() proto.CountPriceServiceServer {
	return &countPriceGRPC{
		db: config.GetDB(),
	}
}
