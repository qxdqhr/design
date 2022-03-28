package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"milkTea/common"
)

type Order struct {
	gorm.Model
	UserId string `json:"userid"`
	CustomerId    string `json:"customerid"`
	OrderingTime string `json:"orderingtime"`   //下单时间
	JuiceNumber string `json:"juicenumber"`     //本单饮品数量
	TotalSellingPrice string `json:"totalsellingprice"`//总售价
	CurEvaluate string  `json:"curevaluate"`    //本单评价
	Buyingjuice string `json:"buyingjuice"`
}
func AddOrderInfo(order *Order)(error){
	db:=common.GetDB()
	db=db.AutoMigrate(&Order{})

	db.Create(order)
	if err := db.Error; err!=nil {
		fmt.Println(err)
		return err
	}

	return nil
}
func RefreshOrderInfo(userid string) ([]Order,error){
	db:=common.GetDB()
	db=db.AutoMigrate(&Order{})
	orders := []Order{}
	dbs:= db.Where("user_id = ?",userid).Find(&orders)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,orders)

	return orders,nil
}


func QueryOrderCustomerName(order *Order) ([]Order,error){

	db:=common.GetDB()
	db=db.AutoMigrate(&Order{})
	orders := []Order{}
	dbs:= db.Debug().Where("user_id = ? AND customer_id = ?",order.UserId,order.CustomerId).Find(&orders)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",orders)

	return orders,nil
}
func QueryOrderBuyingjuice(order *Order) ([]Order,error){

	db:=common.GetDB()
	db=db.AutoMigrate(&Order{})
	orders := []Order{}
	dbs:= db.Where("user_id = ? AND buyingjuice = ?",order.UserId,order.Buyingjuice).Find(&orders)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,orders)

	return orders,nil
}

func QueryOrderOrderingTime(order *Order) ([]Order,error){

	db:=common.GetDB()
	db=db.AutoMigrate(&Order{})
	orders := []Order{}
	dbs:= db.Where("user_id = ? AND ordering_time = ?",order.UserId,order.OrderingTime).Find(&orders)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,orders)

	return orders,nil
}
func QueryOrderJuiceNumber(order *Order) ([]Order,error){

	db:=common.GetDB()
	db=db.AutoMigrate(&Order{})
	orders := []Order{}
	dbs:= db.Where("user_id = ? AND juice_number = ?",order.UserId,order.JuiceNumber).Find(&orders)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,orders)

	return orders,nil
}

func QueryOrderTotalSellingPrice(order *Order) ([]Order,error){

	db:=common.GetDB()
	db=db.AutoMigrate(&Order{})
	orders := []Order{}
	dbs:= db.Where("user_id = ? AND total_selling_price = ?",order.UserId,order.TotalSellingPrice).Find(&orders)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,orders)

	return orders,nil
}

func QueryOrderCurEvaluate(order *Order) ([]Order,error){

	db:=common.GetDB()
	db=db.AutoMigrate(&Order{})
	orders := []Order{}
	dbs:= db.Where("user_id = ? AND cur_evaluate = ?",order.UserId,order.CurEvaluate).Find(&orders)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,orders)

	return orders,nil
}
