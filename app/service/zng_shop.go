package service

import (
	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type ZngShopService struct {
	
}
type ZngShopServiceInterface interface {
	List(req *request.ZngShopRequest) (*[]model.ZngShop,error)
	Detail(req *request.ZngShopRequest) (*model.ZngShop,error)
	Delete(req *request.ZngShopRequest) (err error)
	Add(req *request.ZngShopRequest) (err error)
	Edit(req *request.ZngShopRequest) (err error)
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc
 */
func ( *ZngShopService ) InitModelDB(req *request.ZngShopRequest)(tx *gorm.DB) {
	db := pkg.MysqlConn.Model(&model.ZngShop{})
	return req.Common(db)
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 获取列表
 */
func (r *ZngShopService) List(req *request.ZngShopRequest) (*[]model.ZngShop,error) {
	req.Actions=2
	var list []model.ZngShop
	req.Data=&list
	err := r.InitModelDB(req).Error
	return &list,err
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 获取详情
 */
func (r *ZngShopService) Detail(req *request.ZngShopRequest) (*model.ZngShop,error) {
	req.Actions=3
	var one model.ZngShop
	req.Data=&one
	err := r.InitModelDB(req).Error
	return &one,err
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 添加数据
 */
func (r *ZngShopService) Add(req *request.ZngShopRequest) (err error) {
	req.Actions=5
	//添加数据请自行处理
	err = r.InitModelDB(req).Error
	return err
}

/*
*@Author Administrator
*@Date 9/4/2021 10:29
*@desc 修改数据
 */
func (r *ZngShopService) Edit(req *request.ZngShopRequest) (err error){
	req.Actions=1
	//req.Data= [...] 修改数据请自行处理好
	err = r.InitModelDB(req).Error
	return err
}
/*
*@Author Administrator
*@Date 9/4/2021 10:29
*@desc 修改数据
 */
func (r *ZngShopService) Delete(req *request.ZngShopRequest) (err error){
	req.Actions=4
	err = r.InitModelDB(req).Error
	return  err
}
/*
*@Author Administrator
*@Date 8/4/2021 16:05
*@desc 实例话数据
 */
func NewZngShopService() ZngShopServiceInterface {
	return new(ZngShopService)
}
