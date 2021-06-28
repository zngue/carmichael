package model

import (
	"time"

	"gorm.io/gorm"
)

type ZngUser struct {
	Id          int64          `gorm:"primary_key;auto_increment;column:id;" form:"id" json:"id"  `
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;" form:"deletedAt" json:"deletedAt"  `
	CreatedAt   time.Time      `gorm:"column:created_at;" form:"createdAt" json:"createdAt" `
	UpdatedAt   time.Time      `gorm:"column:updated_at;" form:"updatedAt" json:"updatedAt"  `
	Name        string         `gorm:"column:name;" form:"name" json:"name" `                                      //名称
	NickName    string         `gorm:"column:nickname;" form:"nickname" json:"nickname"  binding:"required"`       //昵称
	Openid      string         `gorm:"column:openid;" form:"openid" json:"openid"  binding:"required"`             //昵称
	HeadImgUrl  string         `gorm:"column:headimgurl;" form:"headimgurl" json:"headimgurl"  binding:"required"` //头像
	ExpireTime  int64          `gorm:"column:expire_time;" form:"expire_time" json:"expire_time"  `                //到期时间
	Account     string         `gorm:"column:account;" form:"account" json:"account" `                             //账号
	ExtAccount  string         `gorm:"column:ext_account;" form:"ext_account" json:"ext_account" `                 //备用账号
	Password    string         `gorm:"column:password;" form:"password" json:"password" `                          //密码
	ExtPassword string         `gorm:"column:ext_password;" form:"ext_password" json:"ext_password" `              //备用密码
	Status      int32          `gorm:"column:status;" form:"status" json:"status"`                                 //  2禁用 1正常
}

func (m *ZngUser) TableName() string {
	return "zng_user"
}
