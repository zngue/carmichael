package request

import (
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type ZngUserRequest struct {
	pkg.CommonRequest
	ID     int    `form:"id" field:"id" where:"eq" default:"0"`
	Openid string `form:"openid" field:"openid" where:"eq" default:""`
	Status string `form:"status" field:"status" where:"eq" default:"0"`
}

func (r *ZngUserRequest) Common(db *gorm.DB) *gorm.DB {
	tx := r.Init(db, *r)
	return tx
}
