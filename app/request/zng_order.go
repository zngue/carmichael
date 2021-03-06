package request

import (
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type ZngOrderRequest struct {
	pkg.CommonRequest
	ID          int    `form:"id" field:"id" where:"eq" default:"0"`
	OpenID      string `form:"openid" field:"openid" where:"eq" default:""`
	IsSellAfter int    `form:"isSellAfter" field:"is_sell_after" where:"eq" default:"0"`
	OrderNum    string `form:"order_num" field:"order_num" where:"eq" default:""`
}

func (r *ZngOrderRequest) Common(db *gorm.DB) *gorm.DB {
	if r.Actions == 2 {
		r.OrderMap = map[string]interface{}{
			"1": "id desc",
		}
		r.OrderString = "1"
	}
	tx := r.Init(db, *r)
	return tx
}
