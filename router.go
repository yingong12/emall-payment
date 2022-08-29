package main

import (
	"emall/http/handler"

	"github.com/gin-gonic/gin"
)

// 绑定routes
func registerRouters() (r *gin.Engine) {
	r = gin.New()
	r.GET("ping", func(c *gin.Context) {
		c.Writer.WriteString("pong")
	})
	//
	order := r.Group("order")
	{
		order.POST("place", handler.PlaceOrder)
		order.POST("pay", handler.PayOrder)
	}
	r.GET("export_data", handler.Export)
	return
}
