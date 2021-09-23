package service

import (
	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type ZngOrderService struct {
}
type ZngOrderServiceInterface interface {
	List(req *request.ZngOrderRequest) (*[]model.ZngOrder, error)
	Detail(req *request.ZngOrderRequest) (*model.ZngOrder, error)
	Delete(req *request.ZngOrderRequest) (err error)
	Add(req *request.ZngOrderRequest) (err error)
	Edit(req *request.ZngOrderRequest) (err error)
	Count(openid string, shopId int64) (int64, error)
}

/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc
 */
func (*ZngOrderService) InitModelDB(req *request.ZngOrderRequest) (tx *gorm.DB) {
	db := pkg.MysqlConn.Model(&model.ZngOrder{})
	return req.Common(db)
}

/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 获取列表
 */
func (r *ZngOrderService) List(req *request.ZngOrderRequest) (*[]model.ZngOrder, error) {
	req.Actions = 2
	var list []model.ZngOrder
	req.Data = &list
	err := r.InitModelDB(req).Error
	return &list, err
}

/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 获取详情
 */
func (r *ZngOrderService) Detail(req *request.ZngOrderRequest) (*model.ZngOrder, error) {
	req.Actions = 3
	var one model.ZngOrder
	req.Data = &one
	err := r.InitModelDB(req).Error
	return &one, err
}

/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 添加数据
 */
func (r *ZngOrderService) Add(req *request.ZngOrderRequest) (err error) {
	req.Actions = 5
	//添加数据请自行处理
	err = r.InitModelDB(req).Error
	return err
}

/*
*@Author Administrator
*@Date 9/4/2021 10:29
*@desc 修改数据
 */
func (r *ZngOrderService) Edit(req *request.ZngOrderRequest) (err error) {
	req.Actions = 1
	//req.Data= [...] 修改数据请自行处理好
	err = r.InitModelDB(req).Error
	return err
}

/*
*@Author Administrator
*@Date 9/4/2021 10:29
*@desc 修改数据
 */
func (r *ZngOrderService) Delete(req *request.ZngOrderRequest) (err error) {
	req.Actions = 4
	err = r.InitModelDB(req).Error
	return err
}
func (r *ZngOrderService) Count(openid string, shopId int64) (int64, error) {
	var num int64
	err := pkg.MysqlConn.Model(&model.ZngOrder{}).Where("shop_id = ?", shopId).Where("pay_status = ?", 1).Where("openid = ?", openid).Count(&num).Error
	return num, err
}

/*
*@Author Administrator
*@Date 8/4/2021 16:05
*@desc 实例话数据
 */
func NewZngOrderService() ZngOrderServiceInterface {
	return new(ZngOrderService)
}
