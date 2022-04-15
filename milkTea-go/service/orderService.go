package service

import (
	"fmt"
	"milkTea/common"
	"strconv"
)

type Order struct {
	Id string `json:"orderid"`
	UserId string `json:"user_id"`
	CustomerId    string `json:"customerid"`
	OrderingTime string `json:"orderingtime"`   //下单时间
	JuiceNumber string `json:"juicenumber"`     //本单饮品数量
	TotalSellingPrice string `json:"totalsellingprice"`//总售价
	CurEvaluate string  `json:"curevaluate"`    //本单评价
	Buyingjuice string `json:"buyingjuice"`
}

func GetInexpenseOfOrder(userid string)(map[string]string, map[string]map[string]string,error)  {
	//先查所有order
	db:=common.GetDB()
	db.AutoMigrate(&Order{})
	orders := []Order{}
	dbs:= db.Where("user_id = ?",userid).Find(&orders)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,nil,err

	}
	monthPriceMapper := make(map[string]string,0)
	for _,val:= range orders {
		month := val.OrderingTime[0:7]
		price := val.TotalSellingPrice
		originPrice,ok := monthPriceMapper[month]
		if ok{
			//存在这个月份,计算收益
			originPrice,err := strconv.ParseFloat(originPrice,64)
			orderPrice,err := strconv.ParseFloat(price,64)
			fmt.Println(originPrice,orderPrice)
			monthPriceMapper[month] = strconv.FormatFloat(originPrice+orderPrice,'f',2,64)
			if err!= nil {
				fmt.Println("atoi错误1",err)
				return nil,nil,fmt.Errorf("atoi错误1")
			}
		}else{
			valPrice,err := strconv.ParseFloat(price,64)
			if err!= nil {
				fmt.Println("atoi错误1",err)
				return nil,nil,fmt.Errorf("atoi错误1")
			}
			monthPriceMapper[month] = strconv.FormatFloat(valPrice,'f',2,64)
		}
	}
	monthTOjuiceNumMapper := make(map[string]map[string]string,0)

	for month,_ := range monthPriceMapper{
		_,ok := monthTOjuiceNumMapper[month]
		if !ok {
			juiceNumMapper := make(map[string]string,0)
			juiceNumMapper["珍珠奶茶"] = "0.00"
			juiceNumMapper["经典奶茶"] = "0.00"
			juiceNumMapper["经典青汁"] = "0.00"
			juiceNumMapper["经典果茶"] = "0.00"
			monthTOjuiceNumMapper[month] = juiceNumMapper
		}
		for _,val:= range orders {
			monthOrder := val.OrderingTime[0:7]
			if month == monthOrder{
				for name,num := range parseJuiceInfo(&val){
					originNum,err1 := strconv.ParseFloat(monthTOjuiceNumMapper[month][name],64)
					curNum,err2 := strconv.ParseFloat(num,64)
					if err1 != nil || err2 != nil {
						fmt.Println("atoi错误2",err1,err2)
						return nil,nil,fmt.Errorf("atoi错误2")
					}
					monthTOjuiceNumMapper[month][name] =
						strconv.FormatFloat(curNum + originNum,'f',2,64)
				}
			}
		}
	}



	return monthPriceMapper,monthTOjuiceNumMapper,nil
}
func AddOrderInfo(order *Order)(error){
	db:=common.GetDB()
	db.AutoMigrate(&Order{})

	dbc := db.Debug().Create(order)
	if err := dbc.Error; err!=nil || dbc.RowsAffected < 1{
		fmt.Println(err)
		return err
	}

	return nil
}
func RefreshOrderInfo(userid string) ([]Order,error){
	db:=common.GetDB()
	db.AutoMigrate(&Order{})
	orders := []Order{}
	dbs:= db.Debug().Where("user_id = ?",userid).Find(&orders)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,orders)

	return orders,nil
}


func QueryOrderCustomerName(order *Order) ([]Order,error){

	db:=common.GetDB()
	db.AutoMigrate(&Order{})
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
	db.AutoMigrate(&Order{})
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
	db.AutoMigrate(&Order{})
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
	db.AutoMigrate(&Order{})
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
	db.AutoMigrate(&Order{})
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
	db.AutoMigrate(&Order{})
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
	db.AutoMigrate(&Order{})
	o := Order{}
	dbs:= db.Where("id = ? ",order.Id).Find(&o)

	if err := dbs.Error; err!=nil || dbs.RowsAffected == 0 {
		fmt.Println(err)
		return err
	}
	//查到了，更新数据
	dbu := db.Debug().Where("id = ?", order.Id).Updates(Order{
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
	db.AutoMigrate(&Order{})
	o := Order{}
	dbs:= db.Where("id = ? ",order.Id).Find(&o)

	if err := dbs.Error; err!=nil || dbs.RowsAffected == 0 {
		fmt.Println(err)
		return err
	}
	//查到了，更新数据
	dbu := db.Where("id = ?", order.Id).Delete(&Order{})
	if err := dbu.Error; err!=nil || dbu.RowsAffected <= 0 {
		fmt.Println(err)
		return err
	}
	return nil
}
