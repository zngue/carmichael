package model

import (
	"time"

	"gorm.io/gorm"
)

type ZngOrder struct {
	Id              int64          `gorm:"primary_key;auto_increment;column:id;" form:"id" json:"id"  binding:""`
	OrderNum        string         `gorm:"column:order_num;" form:"orderNum" json:"orderNum"  binding:"required"`                       //订单编号
	ShopTitle       string         `gorm:"column:order_title;" form:"orderTitle" json:"orderTitle"  binding:"required"`                 //订单编号
	OpenID          string         `gorm:"column:openid;" form:"openid" json:"openid"  `                                                //订单编号
	ShopId          int64          `gorm:"column:shop_id;" form:"shopId" json:"shopId"  binding:"required"`                             //商品id
	UserId          int64          `gorm:"column:user_id;" form:"userId" json:"userId"  binding:"required"`                             //用户id
	ShopTotailPrice float64        `gorm:"column:shop_totail_price;" form:"shopTotailPrice" json:"shopTotailPrice"  binding:"required"` //商品总价
	AddTimeRand     int64          `gorm:"column:add_time_rand;" form:"addTimeRand" json:"addTimeRand"  binding:"required"`
	ShopPrice       float64        `gorm:"column:shop_price;" form:"shopPrice" json:"shopPrice"  binding:"required"`       //商品价格
	ShopNum         int64          `gorm:"column:shop_num;" form:"shopNum" json:"shopNum"  binding:"required"`             //购买商品数量
	AddTime         int64          `gorm:"column:add_time;" form:"addTime" json:"addTime"  binding:"required"`             //创建时间
	UpdateTime      int64          `gorm:"column:update_time;" form:"updateTime" json:"updateTime"  binding:"required"`    //更新时间
	OrderStatus     int32          `gorm:"column:order_status;" form:"orderStatus" json:"orderStatus"  binding:"required"` //1正常 2 删除
	PayStatus       int32          `gorm:"column:pay_status;" form:"payStatus" json:"payStatus"  binding:"required"`       //0 未支付 1 已支付  3 支付超时
	IsSellAfter     int32          `gorm:"column:is_sell_after;" form:"is_sell_after" json:"isSellAfter" `                 //售后处理  1 正常  2 正在售后 3 已售后
	UpdatedAt       time.Time      `gorm:"column:updated_at;" form:"updatedAt" json:"updatedAt"  binding:"required"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;" form:"deletedAt" json:"deletedAt"  binding:"required"`
	CreatedAt       time.Time      `gorm:"column:created_at;" form:"createdAt" json:"createdAt"  binding:"required"`
	UserName        string         `gorm:"column:username;" form:"username" json:"username"  `
	Telphone        string         `gorm:"column:telphone;" form:"telphone" json:"telphone"`
	Address         string         `gorm:"column:address;" form:"address" json:"address"  `
	Desc            string         `gorm:"column:desc;" form:"desc" json:"desc" `
}

func (m *ZngOrder) TableName() string {
	return "zng_order"
}
