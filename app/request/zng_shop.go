package request

import (
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type ZngShopRequest struct {
	pkg.CommonRequest
	ID     int `form:"id" field:"id" where:"eq" default:"0"`
	Status int `form:"status" field:"status" where:"eq" default:"0"`
}

func (r *ZngShopRequest) Common(db *gorm.DB) *gorm.DB {
	tx := r.Init(db, *r)
	return tx
}
