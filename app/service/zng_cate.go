package service

import (
	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type ZngCateService struct {
	
}
type ZngCateServiceInterface interface {
	List(req *request.ZngCateRequest) (*[]model.ZngCate,error)
	Detail(req *request.ZngCateRequest) (*model.ZngCate,error)
	Delete(req *request.ZngCateRequest) (err error)
	Add(req *request.ZngCateRequest) (err error)
	Edit(req *request.ZngCateRequest) (err error)
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc
 */
func ( *ZngCateService ) InitModelDB(req *request.ZngCateRequest)(tx *gorm.DB) {
	db := pkg.MysqlConn.Model(&model.ZngCate{})
	return req.Common(db)
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 获取列表
 */
func (r *ZngCateService) List(req *request.ZngCateRequest) (*[]model.ZngCate,error) {
	req.Actions=2
	var list []model.ZngCate
	req.Data=&list
	err := r.InitModelDB(req).Error
	return &list,err
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 获取详情
 */
func (r *ZngCateService) Detail(req *request.ZngCateRequest) (*model.ZngCate,error) {
	req.Actions=3
	var one model.ZngCate
	req.Data=&one
	err := r.InitModelDB(req).Error
	return &one,err
}
/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 添加数据
 */
func (r *ZngCateService) Add(req *request.ZngCateRequest) (err error) {
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
func (r *ZngCateService) Edit(req *request.ZngCateRequest) (err error){
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
func (r *ZngCateService) Delete(req *request.ZngCateRequest) (err error){
	req.Actions=4
	err = r.InitModelDB(req).Error
	return  err
}
/*
*@Author Administrator
*@Date 8/4/2021 16:05
*@desc 实例话数据
 */
func NewZngCateService() ZngCateServiceInterface {
	return new(ZngCateService)
}
