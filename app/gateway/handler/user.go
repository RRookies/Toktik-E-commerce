package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

func UserRouter(c *route.RouterGroup) {
	c.POST("/info", userInfo)
}

func userInfo(c context.Context, ctx *app.RequestContext) {
	// TODO
}
