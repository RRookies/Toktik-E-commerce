package handler

import (
	"Tiktok/app/order/global"
	"Tiktok/app/order/idl/gen"
	"Tiktok/app/order/model"
	"context"
	"strconv"

	"google.golang.org/grpc/status"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

type OrderService struct {
	gen.UnimplementedOrderServiceServer
}

func (*OrderService) ListOrder(_ context.Context, req *gen.ListOrderReq) (*gen.ListOrderResp, error){
	var orders []model.Order
	var rsp gen.ListOrderResp
	var err error

	orders,err = model.GetOrderList(global.DB,int64(req.UserId),int(req.Pages),int(req.Pagesize))
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

func (*OrderService) PlaceOrder(_ context.Context, req *gen.PlaceOrderReq) (*gen.PlaceOrderResp, error){
	var rsp gen.PlaceOrderResp
	var err error

    for _, item := range req.Items {
        goodsID := strconv.FormatUint(item.ProductsId, 10)
        productrsp,err := global.ProductsSevClient.GetGoodsDetail(context.Background(), &gen.GoodsDetailRequest{
            GoodsId: goodsID,
        })

        if err != nil {
            zap.S().Errorf("调用查询商品detail服务失败: %v", err)
            return nil, status.Errorf(codes.Internal, "服务暂不可用")
        }

        if productrsp == nil  {
            return nil, status.Errorf(codes.NotFound,"商品 %d 不存在", item.ProductsId)
        } else if productrsp.TotalCount < item.Quantity {
            return nil, status.Errorf(
                codes.ResourceExhausted,
                "商品 %d 库存不足 (剩余%d，需要%d)",
                item.ProductsId, productrsp.TotalCount, item.Quantity,
            )
        }
    }

	rsp.Order.OrderId,err = model.CreateOrder(global.DB,req)
    if err != nil {
        return nil, err
    }

    return &rsp,nil
}

func CancelOrder(_ context.Context, req *gen.CancelOrderReq) (*gen.CancelOrderResp, error){
	var rsp gen.CancelOrderResp
	var err error

    err = model.CancelOrder(global.DB,uint64(req.OrderId))
    if err!= nil {
        rsp.Success = false
    }
    return &rsp,nil
}