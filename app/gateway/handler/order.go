package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

func OrderRouter(c *route.RouterGroup) {
	c.POST("/place", orderPlace)
	c.POST("/checkout", orderCheckout)
	c.POST("/list", orderList)
	c.POST("/cancel", orderCancel)
}

func orderPlace(c context.Context, ctx *app.RequestContext) {
	// TODO
}

func orderCheckout(c context.Context, ctx *app.RequestContext) {
	// TODO
}

func orderList(c context.Context, ctx *app.RequestContext) {
	// TODO
}

func orderCancel(c context.Context, ctx *app.RequestContext) {
	// TODO
}
