package service

import (
	"Toktik/app/goods/biz/model"
	"Toktik/app/goods/idl/goods"
	"context"
	"testing"
	"time"

	"Toktik/app/goods/biz/service"

	"github.com/stretchr/testify/suite"
)

type GoodsDetailServiceTestSuite struct {
	BaseTestSuite
	Service *service.GoodsDetailService
}

func (suite *GoodsDetailServiceTestSuite) SetupSuite() {
	// 调用父类的 SetupSuite 初始化数据库
	suite.BaseTestSuite.SetupSuite()

	// 创建 GoodsDetailService 实例
	suite.Service = service.NewGoodsDetailService(suite.DB)
}

func (suite *GoodsDetailServiceTestSuite) TestGetGoodsDetail() {
	// 插入测试商品数据
	gds := &model.Goods{
		GoodsID:     "1",
		CategoryID:  "1",
		Name:        "Test Goods",
		Description: "Test Description",
		Price:       100,
		ImageURL:    "http://example.com/image.jpg",
		Stock_num:   10,
		CreateTime:  time.Now(),
	}
	if err := suite.DB.Create(gds).Error; err != nil {
		suite.T().Fatalf("failed to create test goods: %v", err)
	}

	// 调用 GetGoodsDetail 方法获取商品详情
	req := &goods.GoodsDetailRequest{
		GoodsId: "1",
	}
	goodsDetail, err := suite.Service.GetGoodsDetail(context.Background(), req)
	if err != nil {
		suite.T().Fatalf("failed to get goods detail: %v", err)
	}

	// 断言商品详情是否正确
	suite.Equal(gds.GoodsID, goodsDetail.GoodsId)
	suite.Equal(gds.Name, goodsDetail.Name)
	suite.Equal(gds.Description, goodsDetail.Description)
	suite.Equal(gds.Price, goodsDetail.Price)
	suite.Equal(gds.ImageURL, goodsDetail.ImageUrl)
	suite.Equal(int32(gds.Stock_num), goodsDetail.TotalCount)
	suite.Equal(gds.CreateTime.Unix(), goodsDetail.CreateTime.Seconds)
}

func TestGoodsDetailServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GoodsDetailServiceTestSuite))
}
