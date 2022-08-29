package handler

import (
	"emall/http/structs"
	"emall/providers"
	"emall/service"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

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

type ExportStrct struct {
	Name     string   `gorm:"column:name" json:"name"`
	Title    string   `gorm:"worktitle" json:"title"`
	Userid   int      `gorm:"userid" json:"userid"`
	Looks    string   `gorm:"-" json:"looks"`
	Comments []string `gorm:"-" json:"comments"`
	Img      string   `gorm:"-" json:"img"`
	Url      string   `gorm:"-" json:"url"`
}

// 标题、内容、发布账号、浏览量、评论内容、封面图片、url链接
func Export(c *gin.Context) {
	en := []ExportStrct{}
	tx := providers.DBconnector.Raw("select userid, tag_ids, update_time, workid,category ,whole_score, opt_time, songid from zt_audit.audit_result where `opt_time` > '2022-08-20' and  `final_audit_status` = 1 and `category`!=7 and `whole_score` >= 3.0 and `status` = 0").
		Scan(&en)
	if tx.Error != nil {
		panic(tx.Error)
	}
	//获取其他信息
	csvFile, err := os.Create("./data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)

	for _, usance := range en {
		var row []string
		uidstr := strconv.Itoa(usance.Userid)
		row = append(row, usance.Name)
		row = append(row, usance.Title)
		row = append(row, uidstr)
		writer.Write(row)
	}
	// remember to flush!
	writer.Flush()
	c.Writer.Write(([]byte)("ok"))
}
