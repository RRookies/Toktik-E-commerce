package biz

import (
	"Tiktok/app/cart/internal/dao"
	"Tiktok/app/cart/internal/model"
	db "Tiktok/app/cart/pkg/db"
)

// AddItemToCart 业务逻辑：将商品添加到购物车
func AddItemToCart(userID, productID string, quantity int) (*model.CartItem, error) {
	item := &model.CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}
	err := dao.SaveCartItem(db.DB, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// ListCartItems 业务逻辑：列出购物车中的商品
func ListCartItems(userID string) ([]model.CartItem, error) {
	return dao.GetCartItemsByUserID(db.DB, userID)
}

// DeleteItemFromCart 业务逻辑：从购物车中删除商品
func DeleteItemFromCart(userID, productID string) error {
	return dao.DeleteCartItem(db.DB, userID, productID)
}

// ClearUserCart 业务逻辑：清空用户的购物车
func ClearUserCart(userID string) error {
	return dao.ClearCart(db.DB, userID)
}

// UpdateCartItemQty 业务逻辑：更新购物车中商品的数量
func UpdateCartItemQty(userID, productID string, newQuantity int) error {
	return dao.UpdateCartItemQuantity(db.DB, userID, productID, newQuantity)
}
