package server

import (
	"Toktik/app/goods/biz/service"
	"Toktik/app/goods/idl/goods"
	"context"

	"gorm.io/gorm"
)

type GoodsServer struct {
	ctx context.Context
	db  *gorm.DB
	gds *service.GoodsDetailService
	gcs *service.GoodsCategoryService
	gss *service.GoodsSearchService
	goods.UnimplementedGoodsServiceServer
}

func NewGoodsServer(ctx context.Context, db *gorm.DB) *GoodsServer {
	return &GoodsServer{
		ctx: ctx,
		db:  db,
		gds: service.NewGoodsDetailService(db),
		gcs: service.NewGoodsCategoryService(db),
		gss: service.NewGoodsSearchService(db),
	}
}

// GetGoodsDetail 获取商品详情
func (gs *GoodsServer) GetGoodsDetail(ctx context.Context, req *goods.GoodsDetailRequest) (*goods.GoodsDetailResponse, error) {
	// 调用业务层服务
	resp, err := gs.gds.GetGoodsDetail(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetGoodsCategory 获取商品分类
func (gs *GoodsServer) GetGoodsCategory(ctx context.Context, req *goods.GoodsCategoryRequest) (*goods.GoodsCategoryResponse, error) {
	// 调用业务层服务
	resp, err := gs.gcs.GetGoodsCategory(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchGoods 执行商品搜索
func (gs *GoodsServer) SearchGoods(ctx context.Context, req *goods.GoodsSearchRequest) (*goods.GoodsSearchResponse, error) {
	// 调用业务层服务
	resp, err := gs.gss.SearchGoods(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
