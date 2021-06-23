package request
import (
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)
type ZngCateRequest struct {
	pkg.CommonRequest
	ID int `form:"id" field:"id" where:"eq" default:"0"`
}
func (r *ZngCateRequest) Common(db *gorm.DB) *gorm.DB {
	tx := r.Init(db, *r)
	return tx
}