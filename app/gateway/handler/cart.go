package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

func CartRouter(c *route.RouterGroup) {
	c.POST("/add", cartAdd)
	c.POST("/list", cartList)
	c.POST("/remove", cartRemove)
	c.POST("/clear", cartClear)
	c.POST("/quantity", cartQuantity)
}

func cartAdd(c context.Context, ctx *app.RequestContext) {
	// TODO
}

func cartList(c context.Context, ctx *app.RequestContext) {
	// TODO
}

func cartRemove(c context.Context, ctx *app.RequestContext) {
	// TODO
}

func cartClear(c context.Context, ctx *app.RequestContext) {
	// TODO
}

func cartQuantity(c context.Context, ctx *app.RequestContext) {
	// TODO
}
