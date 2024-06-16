package api

import (
	"app/config"
	"app/grpc/proto"
	"app/model"
	"context"
	"log"
	"sync"

	"gorm.io/gorm"
)

type countPriceGRPC struct {
	db             *gorm.DB
	productService proto.ProductServiceClient
	proto.UnsafeCountPriceServiceServer
}

func (g *countPriceGRPC) CountPriceOrder(ctx context.Context, req *proto.CountPriceOrderReq) (*proto.CountPriceOrderRes, error) {
	var errCount error = nil
	var price float64 = 0

	listProductId := []string{}
	for _, item := range req.Orders {
		listProductId = append(listProductId, item.ProductId)
	}

	resListProduct, errListProduct := g.productService.GetProductByListId(
		context.Background(),
		&proto.GetProductByListIdReq{
			ProductIds: listProductId,
		},
	)
	if errListProduct != nil {
		return nil, errListProduct
	}

	mapProduct := make(map[string]*proto.Product)
	for _, item := range resListProduct.Products {
		mapProduct[item.Id] = item
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, item := range req.Orders {
		wg.Add(1)
		go func(order *proto.Order) {
			log.Println(order)
			if order.TypeInWarehouseId == 0 {
				var warehouse *model.Warehouse
				err := g.db.
					Model(&model.Warehouse{}).
					Where("id = ?", order.WarehouseId).
					First(&warehouse).
					Error

				if err != nil {
					errCount = err
					wg.Done()
					return
				}

				if warehouse != nil {
					mutex.Lock()
					price += float64(mapProduct[order.ProductId].Price) * float64(order.Amount)
					mutex.Unlock()
				}
			} else {
				var typeInWarehouse *model.TypeInWarehouse
				err := g.db.
					Model(&model.TypeInWarehouse{}).
					Preload("Warehouse").
					Where("id = ?", order.TypeInWarehouseId).
					First(&typeInWarehouse).
					Error

				if err != nil {
					errCount = err
					wg.Done()
					return
				}

				if typeInWarehouse != nil {
					mutex.Lock()
					if typeInWarehouse.Price != nil {
						price += *typeInWarehouse.Price * float64(order.Amount)
					} else {
						price += float64(mapProduct[order.ProductId].Price) * float64(order.Amount)
					}
					mutex.Unlock()
				}
			}

			wg.Done()
		}(item)
	}

	wg.Wait()

	if errCount != nil {
		return nil, errCount
	}

	res := &proto.CountPriceOrderRes{
		GroupOrderId: req.GroupOrderId,
		Price:        float32(price),
	}

	return res, nil
}

func NewCountPriceGRPC() proto.CountPriceServiceServer {
	return &countPriceGRPC{
		db:             config.GetDB(),
		productService: proto.NewProductServiceClient(config.GetConnProductGRPC()),
	}
}
