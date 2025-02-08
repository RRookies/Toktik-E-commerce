package handler

import (
	"Tiktok/app/cart/idl/cart"
	"Tiktok/app/cart/internal/biz"
	"context"
)

type CartHandler struct {
	cart.UnimplementedCartServiceServer
}

func (h *CartHandler) AddItem(ctx context.Context, req *cart.AddItemRequest) (*cart.AddItemResponse, error) {
	_, err := biz.AddItemToCart(req.UserId, req.ProductId, int(req.Quantity))
	if err != nil {
		return &cart.AddItemResponse{Success: false, Message: err.Error()}, nil
	}
	// 假设 cart.AddItemResponse 有一个 CartItem 字段用于返回添加的商品信息
	return &cart.AddItemResponse{
		Success: true,
		Message: "Item added successfully",
	}, nil
}

func (h *CartHandler) ListItems(ctx context.Context, req *cart.ListItemsRequest) (*cart.ListItemsResponse, error) {
	items, err := biz.ListCartItems(req.UserId)
	if err != nil {
		return &cart.ListItemsResponse{}, err
	}
	respItems := make([]*cart.CartItem, len(items))
	for i, item := range items {
		respItems[i] = &cart.CartItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
		}
	}
	return &cart.ListItemsResponse{Items: respItems}, nil
}

func (h *CartHandler) DeleteItem(ctx context.Context, req *cart.DeleteItemRequest) (*cart.DeleteItemResponse, error) {
	err := biz.DeleteItemFromCart(req.UserId, req.ProductId)
	if err != nil {
		return &cart.DeleteItemResponse{Success: false, Message: err.Error()}, nil
	}
	return &cart.DeleteItemResponse{Success: true, Message: "Item deleted successfully"}, nil
}

func (h *CartHandler) ClearCart(ctx context.Context, req *cart.ClearCartRequest) (*cart.ClearCartResponse, error) {
	err := biz.ClearUserCart(req.UserId)
	if err != nil {
		return &cart.ClearCartResponse{Success: false, Message: err.Error()}, nil
	}
	return &cart.ClearCartResponse{Success: true, Message: "Cart cleared successfully"}, nil
}

func (h *CartHandler) UpdateItemQuantity(ctx context.Context, req *cart.UpdateItemQuantityRequest) (*cart.UpdateItemQuantityResponse, error) {
	err := biz.UpdateCartItemQty(req.UserId, req.ProductId, int(req.NewQuantity))
	if err != nil {
		return &cart.UpdateItemQuantityResponse{Success: false, Message: err.Error()}, nil
	}
	return &cart.UpdateItemQuantityResponse{Success: true, Message: "Item quantity updated successfully"}, nil
}
