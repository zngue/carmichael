package service

import (
	"github.com/zngue/carmichael/app/model"
	"github.com/zngue/carmichael/app/request"
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type ZngUserService struct {
}
type ZngUserServiceInterface interface {
	List(req *request.ZngUserRequest) (*[]model.ZngUser, error)
	Detail(req *request.ZngUserRequest) (*model.ZngUser, error)
	Delete(req *request.ZngUserRequest) (err error)
	Add(req *request.ZngUserRequest) (err error)
	Edit(req *request.ZngUserRequest) (err error)
}

/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc
 */
func (*ZngUserService) InitModelDB(req *request.ZngUserRequest) (tx *gorm.DB) {
	db := pkg.MysqlConn.Model(&model.ZngUser{})
	return req.Common(db)
}

/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 获取列表
 */
func (r *ZngUserService) List(req *request.ZngUserRequest) (*[]model.ZngUser, error) {
	req.Actions = 2
	var list []model.ZngUser
	req.Data = &list
	err := r.InitModelDB(req).Error
	return &list, err
}

/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 获取详情
 */
func (r *ZngUserService) Detail(req *request.ZngUserRequest) (*model.ZngUser, error) {
	req.Actions = 3
	var one model.ZngUser
	req.Data = &one
	err := r.InitModelDB(req).Error
	return &one, err
}

/*
*@Author Administrator
*@Date 9/4/2021 10:04
*@desc 添加数据
 */
func (r *ZngUserService) Add(req *request.ZngUserRequest) (err error) {
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
func (r *ZngUserService) Edit(req *request.ZngUserRequest) (err error) {
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
func (r *ZngUserService) Delete(req *request.ZngUserRequest) (err error) {
	req.Actions = 4
	err = r.InitModelDB(req).Error
	return err
}

/*
*@Author Administrator
*@Date 8/4/2021 16:05
*@desc 实例话数据
 */
func NewZngUserService() ZngUserServiceInterface {
	return new(ZngUserService)
}
