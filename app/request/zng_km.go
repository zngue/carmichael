package request

import (
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type ZngKmRequest struct {
	pkg.CommonRequest
	ID int `form:"id" field:"id" where:"eq" default:"0"`
}

func (r *ZngKmRequest) Common(db *gorm.DB) *gorm.DB {
	tx := r.Init(db, *r)
	return tx
}
