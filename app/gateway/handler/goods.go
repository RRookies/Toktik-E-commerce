package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

func GoodsRouter(c *route.RouterGroup) {
	c.POST("/detail", goodsDetail)
	c.POST("/category", goodsCategory)
	c.POST("/search", goodsSearch)
}

func goodsDetail(c context.Context, ctx *app.RequestContext) {
	// TODO
}

func goodsCategory(c context.Context, ctx *app.RequestContext) {
	// TODO
}

func goodsSearch(c context.Context, ctx *app.RequestContext) {
	// TODO
}
