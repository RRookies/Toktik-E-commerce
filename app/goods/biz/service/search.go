package service

import (
	"Toktik/app/goods/biz/model"
	"Toktik/app/goods/idl/goods"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// GoodsSearchService 执行商品搜索的服务
type GoodsSearchService struct {
	DB *gorm.DB
}

// NewGoodsSearchService 创建 GoodsSearchService 实例
func NewGoodsSearchService(db *gorm.DB) *GoodsSearchService {
	return &GoodsSearchService{DB: db}
}

// SearchGoods 执行商品搜索
func (s *GoodsSearchService) SearchGoods(ctx context.Context, req *goods.GoodsSearchRequest) (*goods.GoodsSearchResponse, error) {
	goodsList, err := model.GetGoodsByName(s.DB, ctx, req.GoodsName)
	if err != nil {
		return nil, err
	}

	var responseGoods []*goods.GoodsDetailResponse
	for _, gds := range goodsList {
		responseGoods = append(responseGoods, &goods.GoodsDetailResponse{
			GoodsId:     gds.GoodsID,
			Name:        gds.Name,
			Description: gds.Description,
			Price:       gds.Price,
			ImageUrl:    gds.ImageURL,
			CreateTime:  timestamppb.New(gds.CreateTime),
		})
	}

	return &goods.GoodsSearchResponse{
		Goods:      responseGoods,
		TotalCount: int32(len(responseGoods)),
	}, nil
}
