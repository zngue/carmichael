package request

import (
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type ZngKmRequest struct {
	pkg.CommonRequest
	ID     int    `form:"id" field:"id" where:"eq" default:"0"`
	ShopID int    `form:"shop_id" field:"openid" where:"eq" default:""`
	Status string `form:"status" field:"status" where:"eq" default:"-1"`
}

func (r *ZngKmRequest) Common(db *gorm.DB) *gorm.DB {
	tx := r.Init(db, *r)
	return tx
}
