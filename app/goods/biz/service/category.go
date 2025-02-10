package service

import (
	"Toktik/app/goods/biz/model"
	"Toktik/app/goods/idl/goods"
	"context"

	"gorm.io/gorm"
)

// GoodsCategoryService 获取商品分类的服务
type GoodsCategoryService struct {
	DB *gorm.DB
}

// NewGoodsCategoryService 创建 GoodsCategoryService 实例
func NewGoodsCategoryService(db *gorm.DB) *GoodsCategoryService {
	return &GoodsCategoryService{DB: db}
}

// GetGoodsCategory 获取商品分类信息
func (s *GoodsCategoryService) GetGoodsCategory(ctx context.Context, req *goods.GoodsCategoryRequest) (*goods.GoodsCategoryResponse, error) {
	goodsList, err := model.GetGoodsCategoryByID(s.DB, ctx, req.CategoryId, int(req.Page), int(req.PageSize)) // 默认每页 10 条数据
	if err != nil {
		return nil, err
	}

	var goodsIDs []string
	for _, goods := range goodsList {
		goodsIDs = append(goodsIDs, goods.GoodsID)
	}

	return &goods.GoodsCategoryResponse{
		CategoryId:   req.CategoryId,
		CategoryName: "Sample Category", // 假设有个固定名称，实际可查询数据库
		GoodsIds:     goodsIDs,
	}, nil
}
