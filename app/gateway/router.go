package main

import (
	"Tiktok/app/gateway/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func registerRouter(c *server.Hertz) {
	user := c.Group("/user")
	handler.UserRouter(user)

	auth := c.Group("/auth")
	handler.AuthRouter(auth)

	cart := c.Group("/cart")
	handler.CartRouter(cart)

	goods := c.Group("/goods")
	handler.GoodsRouter(goods)

	order := c.Group("/order")
	handler.OrderRouter(order)

	pay := c.Group("/pay")
	handler.PayRouter(pay)
}
