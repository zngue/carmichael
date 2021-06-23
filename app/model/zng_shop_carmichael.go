package model


type ZngShopCarmichael struct {
	Id int64 `gorm:"primary_key;auto_increment;column:id;" form:"id" json:"id"  binding:""`
	ShopId int64 `gorm:"column:shop_id;" form:"shopId" json:"shopId"  binding:"required"`  //商品id
	UserId int64 `gorm:"column:user_id;" form:"userId" json:"userId"  binding:"required"`  //卡密
	CarNo string `gorm:"column:car_no;" form:"carNo" json:"carNo"  binding:""`
	Status int32 `gorm:"column:status;" form:"status" json:"status"  binding:"required"`  //默认未使用 1 已经使用
}
func (m *ZngShopCarmichael) TableName() string  {
	return "zng_shop_carmichael"
}