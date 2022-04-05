package service

import (
	"fmt"
	"milkTea/common"
)

type Customer struct {
	Name    string `json:"name"`
	UserId string `json:"userid"`
	CustomerId    string `json:"customerid"`
	BuyingTime string `json:"buying_time"`
	RecentEvaluate string `json:"recent_evaluate"`
}

//添加订单中的顾客信息
func UpdateCustomerInfo(customer *Customer)(error){
	db:=common.GetDB()
	db.AutoMigrate(&Customer{})
	//判断当前uid是否已存在
	c:=new(Customer)
	dbq:=db.First(&c, "customer_id = ?", customer.CustomerId)

	if dbq.RowsAffected<1{//没查到,创建新用户数据
		db.Create(&customer)
		fmt.Println("顾客数据已创建")
	} else if(dbq.RowsAffected == 1){ //有顾客信息，更新顾客信息
		db.Model(&c).Where("customer_id = ? AND user_id = ?", customer.CustomerId,customer.UserId).Updates(map[string]string{
			"buying_time" :customer.BuyingTime,
			"recent_evaluate":customer.RecentEvaluate,
			"user_id":customer.UserId,
		})
		fmt.Println("顾客数据已更新")
	}
	return nil
}

func DeleteCustomerInfo(customer *Customer)(error){
	db:=common.GetDB()
	db.AutoMigrate(&Customer{})
	//判断当前uid是否已存在
	c:=new(Customer)
	dbq:=db.Debug().First(&c, "customer_id = ? AND user_id = ?", customer.CustomerId,customer.UserId)

	if(dbq.RowsAffected == 1){ //有顾客信息，更新顾客信息
		db.Debug().Model(&c).Where("customer_id = ? AND user_id = ?", customer.CustomerId,customer.UserId).Updates(map[string]string{
			"buying_time":"--",
			"recent_evaluate":"--",
		})
	}
	return nil
}
func RefreshCustomerInfo(userid string) ([]Customer,error){
	db:=common.GetDB()
	db.AutoMigrate(&Customer{})
	customers := []Customer{}
	dbs:= db.Debug().Where("user_id = ?",userid).Find(&customers)

	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,customers)

	return customers,nil
}
func QueryCustomerName(customer *Customer) ([]Customer,error) {
	db:=common.GetDB()
	db.AutoMigrate(&Customer{})
	customers := []Customer{}
	dbs:= db.Where("user_id = ? AND name = ?",customer.UserId,customer.Name).Find(&customers)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,orders)

	return customers,nil
}

func QueryCustomerBuyingTime(customer *Customer) ([]Customer,error)  {
	db:=common.GetDB()
	db.AutoMigrate(&Customer{})
	customers := []Customer{}
	dbs:= db.Where("user_id = ? AND buying_time = ?",customer.UserId,customer.BuyingTime).Find(&customers)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,orders)

	return customers,nil
}

func QueryCustomerRecentEvaluate(customer *Customer) ([]Customer,error)  {
	db:=common.GetDB()
	db.AutoMigrate(&Customer{})
	customers := []Customer{}
	dbs:= db.Where("user_id = ? AND recent_evaluate = ?",customer.UserId,customer.RecentEvaluate).Find(&customers)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,orders)

	return customers,nil
}
