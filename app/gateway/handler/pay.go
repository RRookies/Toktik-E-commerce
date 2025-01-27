package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

func PayRouter(c *route.RouterGroup) {
	c.POST("/charge", payCharge)
	c.POST("/callback", payCallback)
}

func payCharge(c context.Context, ctx *app.RequestContext) {
	// TODO
}

func payCallback(c context.Context, ctx *app.RequestContext) {
	// TODO
}
