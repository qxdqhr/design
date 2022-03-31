package service

import (
	"fmt"
	"milkTea/common"
)

type Order struct {
	Id string `json:"orderid"`
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

	dbc := db.Debug().Create(order)
	if err := dbc.Error; err!=nil || dbc.RowsAffected < 1{
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
func ModifyOrderInfo(order *Order) (error){
	db:=common.GetDB()
	db=db.AutoMigrate(&Order{})
	o := Order{}
	dbs:= db.Where("id = ? ",order.Id).Find(&o)

	if err := dbs.Error; err!=nil || dbs.RowsAffected == 0 {
		fmt.Println(err)
		return err
	}
	//查到了，更新数据
	dbu := db.Where("id = ?", order.Id).Update(Order{
		UserId:            order.UserId,
		CustomerId:        order.CustomerId,
		OrderingTime:      order.OrderingTime,
		JuiceNumber:       order.JuiceNumber,
		TotalSellingPrice: order.TotalSellingPrice,
		CurEvaluate:       order.CurEvaluate,
		Buyingjuice:       order.Buyingjuice,
	})
	if err := dbu.Error; err!=nil || dbu.RowsAffected <= 0 {
		fmt.Println(err)
		return err
	}
	return nil
}
func DeleteOrderInfo(order *Order) (error){
	db:=common.GetDB()
	db=db.AutoMigrate(&Order{})
	o := Order{}
	dbs:= db.Where("id = ? ",order.Id).Find(&o)

	if err := dbs.Error; err!=nil || dbs.RowsAffected == 0 {
		fmt.Println(err)
		return err
	}
	//查到了，更新数据
	dbu := db.Where("id = ?", order.Id).Delete(Order{
		UserId:            order.UserId,
		CustomerId:        order.CustomerId,
		OrderingTime:      order.OrderingTime,
		JuiceNumber:       order.JuiceNumber,
		TotalSellingPrice: order.TotalSellingPrice,
		CurEvaluate:       order.CurEvaluate,
		Buyingjuice:       order.Buyingjuice,
	})
	if err := dbu.Error; err!=nil || dbu.RowsAffected <= 0 {
		fmt.Println(err)
		return err
	}
	return nil
}
