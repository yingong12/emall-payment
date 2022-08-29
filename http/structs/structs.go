package structs

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PlaceOrderRequest struct {
	UserID string `json:"userID" binding:"required"`
	SkuID  string `json:"sku_id" binding:"required"`
	Count  int    `json:"count"  binding:"required"`
}
type PlaceOrderResponse struct {
	OrderID string `json:"orderID"`
}

type PayOrderRequest struct {
	OrderID string `json:"orderID"`
}
