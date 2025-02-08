package model

type CartItem struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    string `gorm:"index"`
	ProductID string
	Quantity  int
}
