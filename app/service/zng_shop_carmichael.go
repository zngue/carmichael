package service

import (
	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type ZngShopCarmichaelService struct {
	
}
type ZngShopCarmichaelServiceInterface interface {
	List(req *request.ZngShopCarmichaelRequest) (*[]model.ZngShopCarmichael,error)
	Detail(req *request.ZngShopCarmichaelRequest) (*model.ZngShopCarmichael,error)
	Delete(req *request.ZngShopCarmichaelRequest) (err error)
	Add(req *request.ZngShopCarmichaelRequest) (err error)
	Edit(req *request.ZngShopCarmichaelRequest) (err error)
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc
 */
func ( *ZngShopCarmichaelService ) InitModelDB(req *request.ZngShopCarmichaelRequest)(tx *gorm.DB) {
	db := pkg.MysqlConn.Model(&model.ZngShopCarmichael{})
	return req.Common(db)
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 获取列表
 */
func (r *ZngShopCarmichaelService) List(req *request.ZngShopCarmichaelRequest) (*[]model.ZngShopCarmichael,error) {
	req.Actions=2
	var list []model.ZngShopCarmichael
	req.Data=&list
	err := r.InitModelDB(req).Error
	return &list,err
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 获取详情
 */
func (r *ZngShopCarmichaelService) Detail(req *request.ZngShopCarmichaelRequest) (*model.ZngShopCarmichael,error) {
	req.Actions=3
	var one model.ZngShopCarmichael
	req.Data=&one
	err := r.InitModelDB(req).Error
	return &one,err
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 添加数据
 */
func (r *ZngShopCarmichaelService) Add(req *request.ZngShopCarmichaelRequest) (err error) {
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
func (r *ZngShopCarmichaelService) Edit(req *request.ZngShopCarmichaelRequest) (err error){
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
func (r *ZngShopCarmichaelService) Delete(req *request.ZngShopCarmichaelRequest) (err error){
	req.Actions=4
	err = r.InitModelDB(req).Error
	return  err
}
/*
*@Author Administrator
*@Date 8/4/2021 16:05
*@desc 实例话数据
 */
func NewZngShopCarmichaelService() ZngShopCarmichaelServiceInterface {
	return new(ZngShopCarmichaelService)
}
