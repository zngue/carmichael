package model

import (
	"time"

	"gorm.io/gorm"
)

type ZngCate struct {
	Id         int64          `gorm:"primary_key;auto_increment;column:id;" form:"id" json:"id"  binding:""`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;" form:"deletedAt" json:"deletedAt"  binding:"required"`
	CreatedAt  time.Time      `gorm:"column:created_at;" form:"createdAt" json:"createdAt"  binding:"required"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;" form:"updatedAt" json:"updatedAt"  binding:"required"`
	Name       string         `gorm:"column:name;" form:"name" json:"name"  binding:"required"`        //名称
	ImgUrl     string         `gorm:"column:img_url;" form:"imgUrl" json:"imgUrl"  binding:"required"` //图片地址
	Sort       int64          `gorm:"column:sort;" form:"sort" json:"sort"  binding:"required"`        //排序
	UpdateTime int64          `gorm:"column:update_time;" form:"updateTime" json:"updateTime"  binding:"required"`
	CreateTime int64          `gorm:"column:create_time;" form:"createTime" json:"createTime"  binding:"required"` //创建时间
	Status     int32          `gorm:"column:status;" form:"status" json:"status"  binding:"required"`              //  2禁用 1正常
}

func (m *ZngCate) TableName() string {
	return "zng_cate"
}
