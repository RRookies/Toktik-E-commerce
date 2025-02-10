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

type GoodsCategoryServiceTestSuite struct {
	BaseTestSuite
	Service *service.GoodsCategoryService
}

func (suite *GoodsCategoryServiceTestSuite) SetupSuite() {

	suite.BaseTestSuite.SetupSuite()

	suite.Service = service.NewGoodsCategoryService(suite.DB)
}

func (suite *GoodsCategoryServiceTestSuite) TestGetGoodsCategory() {
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

	req := &goods.GoodsCategoryRequest{
		CategoryId: "1",
		Page:       1,
		PageSize:   10,
	}
	goodsCategory, err := suite.Service.GetGoodsCategory(context.Background(), req)
	if err != nil {
		suite.T().Fatalf("failed to get goods detail: %v", err)
	}
	suite.Equal(gds.CategoryID, goodsCategory.CategoryId)
}

func TestGoodsCategoryServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GoodsCategoryServiceTestSuite))
}
