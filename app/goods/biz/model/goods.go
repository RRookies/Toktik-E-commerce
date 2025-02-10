package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Goods struct {
	*gorm.Model
	GoodsID     string    `gorm:"column:goods_id;not null"`
	CategoryID  string    `gorm:"column:category_id;not null"`
	Name        string    `gorm:"column:name;not null"`
	Description string    `gorm:"column:description;"`
	Price       float32   `gorm:"column:price;not null"`
	ImageURL    string    `gorm:"column:image_url;"`
	Stock_num   int       `gorm:"column:stock_num;not null"`
	CreateTime  time.Time `gorm:"column:create_time;"`
}

// CreateGoods 创建商品
func CreateGoods(db *gorm.DB, ctx context.Context, goods *Goods) error {
	return db.WithContext(ctx).Create(goods).Error
}

// GetGoodsByID 根据商品ID获取商品信息
func GetGoodsByID(db *gorm.DB, ctx context.Context, goodsID string) (*Goods, error) {
	var goods Goods
	err := db.WithContext(ctx).Model(&Goods{}).Where("goods_id = ?", goodsID).First(&goods).Error
	if err != nil {
		return nil, err
	}
	return &goods, nil
}

// GetGoodsByName 根据商品名称获取商品信息
func GetGoodsByName(db *gorm.DB, ctx context.Context, name string) ([]Goods, error) {
	var goodsList []Goods
	err := db.WithContext(ctx).Model(&Goods{}).Where("name LIKE ?", "%"+name+"%").Find(&goodsList).Error
	if err != nil {
		return nil, err
	}
	return goodsList, nil
}

// GetGoodsCategoryByID 根据分类ID获取商品分类下的商品信息
func GetGoodsCategoryByID(db *gorm.DB, ctx context.Context, categoryID string, page, pageSize int) ([]Goods, error) {
	var goodsList []Goods
	offset := (page - 1) * pageSize
	err := db.WithContext(ctx).Model(&Goods{}).Where("category_id = ?", categoryID).
		Limit(pageSize).Offset(offset).Find(&goodsList).Error
	if err != nil {
		return nil, err
	}
	return goodsList, nil
}
