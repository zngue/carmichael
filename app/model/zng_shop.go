package model

import (
	"time"

	"gorm.io/gorm"
)

type ZngShop struct {
	Id          int64          `gorm:"primary_key;auto_increment;column:id;" form:"id" json:"id"  binding:""`
	Title       string         `gorm:"column:title;" form:"title" json:"title"  binding:"required"`     //商品标题
	ItemNo      string         `gorm:"column:item_no;" form:"itemNo" json:"itemNo"  binding:"required"` //商品编号
	CateId      int64          `gorm:"column:cate_id;" form:"cateId" json:"cateId"  binding:"required"` //商品分类id
	AddTimeRand int64          `gorm:"column:add_time_rand;" form:"addTimeRand" json:"addTimeRand"  binding:"required"`
	SalesNum    int64          `gorm:"column:sales_num;" form:"salesNum" json:"salesNum"  binding:"required"`    //商品出售数量
	UpdateNum   int64          `gorm:"column:update_num;" form:"updateNum" json:"updateNum"  binding:"required"` //更新时库存预警值
	ItemNum     int64          `gorm:"column:item_num;" form:"itemNum" json:"itemNum"  binding:"required"`       //商品数量
	Money       int64          `gorm:"column:money;" form:"money" json:"money"  binding:"required"`              //购买金额 单位分
	UserLimit   int64          `gorm:"column:user_limit;" form:"userLimit" json:"userLimit"  binding:"required"` //用户限购
	Img         string         `gorm:"column:img;" form:"img" json:"img"  binding:"required"`                    //轮播图 多张图片用|分割
	Content     string         `gorm:"column:content;" form:"content" json:"content"  binding:"required"`        //详情
	Status      int32          `gorm:"column:status;" form:"status" json:"status"  binding:"required"`           // 1 上架 2 下架  3 草稿 4 删除
	CreatedAt   time.Time      `gorm:"column:created_at;" form:"createdAt" json:"createdAt"  binding:"required"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;" form:"updatedAt" json:"updatedAt"  binding:"required"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;" form:"deletedAt" json:"deletedAt"  binding:"required"`
	ShopSort    int64          `gorm:"column:shop_sort;" form:"shopSort" json:"shopSort"  binding:"required"`
}

func (m *ZngShop) TableName() string {
	return "zng_shop"
}
