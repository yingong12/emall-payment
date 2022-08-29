package models

type OrderModel struct {
	OrderID      string `gorm:"column:order_id"`
	UserID       string `gorm:"column:user_id"`
	SkuID        string `gorm:"column:sku_id"`
	Details      string `gorm:"column:details"`
	State        int    `gorm:"column:state"`
	CreatedAtFmt string `gorm:"-" json:"created_at"` //返回给业务侧
	UpdatedAtFmt string `gorm:"-" json:"udated_at"`
}
