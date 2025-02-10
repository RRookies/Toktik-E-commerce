package service

import (
	"Toktik/app/goods/biz/model"
	"Toktik/app/goods/biz/service"
	"Toktik/app/goods/idl/goods"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type GoodsSearchServiceTestSuite struct {
	BaseTestSuite
	Service *service.GoodsSearchService
}

func (suite *GoodsSearchServiceTestSuite) SetupSuite() {
	// 调用父类的 SetupSuite 初始化数据库
	suite.BaseTestSuite.SetupSuite()

	// 创建 GoodsSearchService 实例
	suite.Service = service.NewGoodsSearchService(suite.DB)
}

func (suite *GoodsSearchServiceTestSuite) TestSearchGoods() {
	// 插入测试商品数据
	gds := &model.Goods{
		GoodsID:     "1",
		CategoryID:  "1",
		Name:        "Test Goods",
		Description: "Test Description",
		Price:       100,
		ImageURL:    "http://example.com/image.jpg",
		CreateTime:  time.Now(),
	}
	if err := suite.DB.Create(gds).Error; err != nil {
		suite.T().Fatalf("failed to create test goods: %v", err)
	}
	req := &goods.GoodsSearchRequest{
		GoodsName: "Test Goods",
	}
	goodsSearch, err := suite.Service.SearchGoods(context.Background(), req)
	if err != nil {
		suite.T().Fatalf("failed to get goods detail: %v", err)
	}
	suite.Equal(gds.GoodsID, goodsSearch.Goods[0].GoodsId)

}

func TestGoodsSearchServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GoodsSearchServiceTestSuite))
}
