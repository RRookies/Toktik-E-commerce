package handler

import (
	"Tiktok/app/order/idl/gen"
	"Tiktok/app/order/model"
	"context"

	"go.uber.org/zap"
)

type OrderService struct {
	gen.UnimplementedOrderServiceServer
}

func (*OrderService) ListOrder(_ context.Context, req *gen.ListOrderReq) (*gen.ListOrderResp, error){
	var orders []model.Order
	var rsp gen.ListOrderResp
	var err error

	orders,err = model.GetOrderList(int64(req.UserId),int(req.Pages),int(req.Pagesize))
	if err != nil {
		zap.S().Error(err)
		return nil,err
	}
	
	for _,order := range orders {
		genOrder := &gen.Order{
            OrderId:     uint64(order.OrderID),
            OrderNo:     order.OrderNo,
            UserId:      uint64(order.UserID),
            TotalPrice:  int32(order.TotalPrice),
            PayStatus:   int32(order.PayStatus),
            PayType:     int32(order.PayType),
            PayTime:     order.PayTime.String(), // 假设 PayTime 是 time.Time 类型
            OrderStatus: int32(order.OrderStatus),
            ExtraInfo:   order.ExtraInfo,
            IsDeleted:   order.IsDeleted,
        }

        // 处理 Address 字段
        if order.Address != nil {
            genOrder.Address = &gen.Address{
                ProvinceName:  order.Address.ProvinceName,
                CityName:      order.Address.CityName,
                RegionName:    order.Address.RegionName,
                DetailAddress: order.Address.DetailAddress,
            }
        }

        // 处理 OrderItem 字段
        for _, item := range order.OrderItems {
            genOrder.OrderItems = append(genOrder.OrderItems, &gen.OrderItem{
                ProductsId :       uint64(item.GoodsID),
                ProductsPrice:  int32(item.SellingPrice),
                Quantity:    int32(item.GoodsCount),
            })
        }

        rsp.Orders = append(rsp.Orders, genOrder)
	}
	return &rsp,nil
}


