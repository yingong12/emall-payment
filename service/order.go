package service

import (
	"emall/http/structs"
	"emall/models"
	"emall/providers"
	"emall/task"
	"emall/utils"
	"encoding/json"
	"log"
)

const RedisScript = `
local hKey = KEYS[1]
local hField = KEYS[2]
local count = ARGV[1]
local result = redis.call('hget',hKey,hField)
if result -count >= 0 then
	redis.call('HINCRBY',hKey,hField,-count)
	return 0
end
return 1
`

func PlaceOrder(req *structs.PlaceOrderRequest) (rsp *structs.PlaceOrderResponse, err error) {
	//
	redisKey := "sku:" + req.SkuID
	redisField := "count"
	count := req.Count
	res := providers.RedisConnector.Eval(RedisScript, []string{redisKey, redisField}, count)
	result, err := res.Result()
	if err != nil {
		log.Fatal(err)
	}
	//超卖
	if result.(int64) == 1 {
		return
	}
	//生成order
	ordrID := utils.GenerateOrderID()
	err = createOrder(ordrID, req.UserID, req.SkuID, count)
	if err != nil {
		return
	}
	//TODO:订单过期逻辑. 时间片轮转
	task.AddTask(ordrID)
	rsp = new(structs.PlaceOrderResponse)
	rsp.OrderID = ordrID
	return
}

// createOrder 新建订单
func createOrder(orderID, userID, skuID string, count int) (err error) {
	details := map[string]interface{}{
		"count": count,
	}
	j, _ := json.Marshal(details)
	en := models.OrderModel{
		UserID:  userID,
		SkuID:   skuID,
		OrderID: orderID,
		Details: string(j),
	}
	//写db
	tx := providers.DBconnector.Table("t_order").
		Create(en)
	return tx.Error
}

// PayOrder 支付
func PayOrder(orderID string) (code int, err error) {
	//扣款
	//订单状态更改。
	return
}
