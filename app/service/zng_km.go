package service

import (
	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type ZngKmService struct {
}
type ZngKmServiceInterface interface {
	List(req *request.ZngKmRequest) (*[]model.ShopKm, error)
	Detail(req *request.ZngKmRequest) (*model.ZngKm, error)
	Delete(req *request.ZngKmRequest) (err error)
	Add(req *request.ZngKmRequest) (err error)
	Edit(req *request.ZngKmRequest) (err error)
}

// InitModelDB /*
func (*ZngKmService) InitModelDB(req *request.ZngKmRequest) (tx *gorm.DB) {
	tx = pkg.MysqlConn
	if req.Actions == 2 {
		tx = tx.Table(new(model.ZngKm).TableName()).Preload("Shop").Order("id desc")
	} else {
		tx = tx.Model(&model.ZngKm{})
	}
	return req.Common(tx)
}

// List /*
func (r *ZngKmService) List(req *request.ZngKmRequest) (*[]model.ShopKm, error) {
	req.Actions = 2
	var list []model.ShopKm
	req.Data = &list
	err := r.InitModelDB(req).Error
	return &list, err
}

// Detail /*
func (r *ZngKmService) Detail(req *request.ZngKmRequest) (*model.ZngKm, error) {
	req.Actions = 3
	var one model.ZngKm
	req.Data = &one
	err := r.InitModelDB(req).Error
	return &one, err
}

// Add /*
func (r *ZngKmService) Add(req *request.ZngKmRequest) (err error) {
	req.Actions = 5
	//添加数据请自行处理
	err = r.InitModelDB(req).Error
	return err
}

// Edit /*
func (r *ZngKmService) Edit(req *request.ZngKmRequest) (err error) {
	req.Actions = 1
	//req.Data= [...] 修改数据请自行处理好
	err = r.InitModelDB(req).Error
	return err
}

// Delete /*
func (r *ZngKmService) Delete(req *request.ZngKmRequest) (err error) {
	req.Actions = 4
	err = r.InitModelDB(req).Error
	return err
}

// NewZngKmService /*
func NewZngKmService() ZngKmServiceInterface {
	return new(ZngKmService)
}
