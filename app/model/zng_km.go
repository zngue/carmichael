package model

import (
	"time"

	"gorm.io/gorm"
)

//卡密
type ZngKm struct {
	Id          int64          `gorm:"primary_key;auto_increment;column:id;" form:"id" json:"id"  `
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;" form:"deletedAt" json:"deletedAt"  `
	CreatedAt   time.Time      `gorm:"column:created_at;" form:"createdAt" json:"createdAt" `
	UpdatedAt   time.Time      `gorm:"column:updated_at;" form:"updatedAt" json:"updatedAt"  `
	ShopID      int64          `gorm:"column:shop_id;" form:"shop_id" json:"shop_id"  `               //商品id
	Account     string         `gorm:"column:account;" form:"account" json:"account" `                //账号
	Password    string         `gorm:"column:password;" form:"password" json:"password" `             //密码
	ExtAccount  string         `gorm:"column:ext_account;" form:"ext_account" json:"ext_account" `    //备用账号
	ExtPassword string         `gorm:"column:ext_password;" form:"ext_password" json:"ext_password" ` //备用密码
	Status      int32          `gorm:"column:status;" form:"status" json:"status"`                    //  0未使用  1已使用 2 作废
	ExpireTime  int64          `gorm:"column:expire_time;" form:"expire_time" json:"expire_time"  `   //有效期
}

func (m *ZngKm) TableName() string {
	return "zng_km"
}
