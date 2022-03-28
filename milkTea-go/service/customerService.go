package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"milkTea/common"
)

type Customer struct {
	gorm.Model
	Name    string `json:"name"`
	CustomerId    string `json:"customerid"`
	BuyingTime string `json:"buying_time"`
	RecentEvaluate string `json:"recent_evaluate"`
}
//添加订单中的顾客信息
func UpdateCustomerInfo(customer *Customer)(error){
	db:=common.GetDB()
	db=db.AutoMigrate(&Customer{})
	//判断当前uid是否已存在
	c:=new(Customer)
	dbq:=db.First(&c, "customer_id = ?", customer.CustomerId)
	defer db.Close()

	if dbq.RowsAffected<1{//没查到,创建新用户数据
		db.Create(customer)
		fmt.Println("顾客数据已创建")
	} else if(dbq.RowsAffected == 1){ //有顾客信息，更新顾客信息
		db.Model(&c).Updates(map[string]string{
			c.BuyingTime :customer.BuyingTime,
			c.RecentEvaluate:customer.RecentEvaluate,
		})
		fmt.Println("顾客数据已更新")
	}
	return nil
}
