package handler

import (
	"emall/http/structs"
	"emall/service"

	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	req := structs.PlaceOrderRequest{}
	if err := c.BindJSON(&req); err != nil {
		return
	}
	res := structs.BaseResponse{Msg: "ok"}
	rsp, err := service.PlaceOrder(&req)
	if err != nil {
		res.Code = -1
		res.Msg = err.Error()
		c.JSON(200, res)
		return
	}
	if rsp == nil {
		res.Code = 1
		res.Msg = "库存不足"
	}
	res.Data = rsp
	c.JSON(200, res)
}

func PayOrder(c *gin.Context) {
	req := structs.PayOrderRequest{}
	if err := c.BindJSON(&req); err != nil {
		return
	}
	res := structs.BaseResponse{Msg: "ok"}
	code, err := service.PayOrder(req.OrderID)
	if err != nil {
		res.Code = -1
		res.Msg = err.Error()
		c.JSON(200, res)
		return
	}
	res.Code = code
	c.JSON(200, res)
	return
}
