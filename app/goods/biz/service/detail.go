package service

import (
	"Toktik/app/goods/biz/model"
	"Toktik/app/goods/idl/goods"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// GoodsDetailService 获取商品详情的服务
type GoodsDetailService struct {
	DB *gorm.DB
}

// NewGoodsDetailService 创建 GoodsDetailService 实例
func NewGoodsDetailService(db *gorm.DB) *GoodsDetailService {
	return &GoodsDetailService{DB: db}
}

// GetGoodsDetail 获取商品详情
func (s *GoodsDetailService) GetGoodsDetail(ctx context.Context, req *goods.GoodsDetailRequest) (*goods.GoodsDetailResponse, error) {
	gds, err := model.GetGoodsByID(s.DB, ctx, req.GoodsId)
	if err != nil {
		return nil, err
	}

	// 返回生成的 GoodsDetailResponse 类型
	return &goods.GoodsDetailResponse{
		GoodsId:     gds.GoodsID,
		Name:        gds.Name,
		Description: gds.Description,
		Price:       gds.Price,
		ImageUrl:    gds.ImageURL,
		TotalCount:  int32(gds.Stock_num),
		CreateTime:  timestamppb.New(gds.CreateTime),
	}, nil
}
