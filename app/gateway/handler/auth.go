package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

func AuthRouter(c *route.RouterGroup) {
	c.POST("/register", authRegister)
	c.POST("/login", authLogin)
	c.POST("/refreshToken", authRefreshToken)
	c.POST("/logout", authLogout)
}

func authRegister(c context.Context, ctx *app.RequestContext) {
	// TODO
}

func authLogin(c context.Context, ctx *app.RequestContext) {
	// TODO
}

func authRefreshToken(c context.Context, ctx *app.RequestContext) {
	// TODO
}

func authLogout(c context.Context, ctx *app.RequestContext) {
	// TODO
}
