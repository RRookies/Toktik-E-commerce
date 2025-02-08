package dao

import (
	"Tiktok/app/cart/internal/model"

	"gorm.io/gorm"
)

// SaveCartItem 保存购物车项到数据库
func SaveCartItem(db *gorm.DB, item *model.CartItem) error {
	return db.Save(item).Error
}

// GetCartItemsByUserID 根据用户 ID 获取购物车中的所有商品项
func GetCartItemsByUserID(db *gorm.DB, userID string) ([]model.CartItem, error) {
	var items []model.CartItem
	result := db.Where("user_id = ?", userID).Find(&items)
	return items, result.Error
}

// DeleteCartItem 根据用户 ID 和商品 ID 删除购物车中的商品项
func DeleteCartItem(db *gorm.DB, userID, productID string) error {
	return db.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&model.CartItem{}).Error
}

// ClearCart 清空指定用户的购物车
func ClearCart(db *gorm.DB, userID string) error {
	return db.Where("user_id = ?", userID).Delete(&model.CartItem{}).Error
}

// UpdateCartItemQuantity 更新购物车中指定商品项的数量
func UpdateCartItemQuantity(db *gorm.DB, userID, productID string, newQuantity int) error {
	return db.Model(&model.CartItem{}).Where("user_id = ? AND product_id = ?", userID, productID).Update("quantity", newQuantity).Error
}
