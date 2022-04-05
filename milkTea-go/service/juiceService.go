package service

import (
	"fmt"
	"milkTea/common"
	"strconv"
	"strings"
)

type Juice struct {
	JuiceId     string `json:"id"`
	JuiceName string `json:"juice_name"`
	JuiceType string `json:"juice_type"`
	LastOrderingTime string `json:"last_ordering_time"`
	Price string `json:"price"`
	Profit string `json:"profit"`
	Cost string `json:"cost"`
	CurEvaluate string `json:"cur_evaluate"` //最近评价
	JuiceSoldNumber string `json:"juice_sold_number"`
	SellingTotalPrice string `json:"selling_total_price"`//历史总销售额
	GoodEvaluateNum string `json:"good_evaluate_num"`//好评数
}
func parseJuiceInfo(order *Order) map[string]string{
	//饮品信息:历史销量,总销售额,好评数
	//aaa:1
	juiceNumMap:= make(map[string]string)
	juices := strings.Split(order.Buyingjuice,"|")

	for  j:=0;j< len(juices);j++{
		juice := strings.Split(juices[j],":")
		juiceNumMap[juice[0]] = juice[1]
	}
	return juiceNumMap
}
func UpdateJuiceWithOrder(order *Order)error{
	db:=common.GetDB()
	db.AutoMigrate(&Juice{})
	for k,v := range parseJuiceInfo(order){
		j:=new(Juice)
		dbq:=db.First(&j, "juice_name = ?", k)
		fmt.Println("查询到此饮品数据：",j)

		if dbq.RowsAffected<1{//没查到
			return fmt.Errorf("无此饮料数据")
		} else if(dbq.RowsAffected == 1){ //有饮品信息，更新饮品信息

			totalPriceNum, err1 := strconv.ParseFloat(j.SellingTotalPrice,64)//当前销售额
			curSoldNum, err2 := strconv.ParseInt(j.JuiceSoldNumber,10,64)//当前已购买数量
			num,err3:=strconv.ParseInt(v,10,64)//饮品数量
			price, err4 := strconv.ParseFloat(j.Price,64)//饮品售价
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				fmt.Println("atoi错误",err1,err2,err3,err4)
				return fmt.Errorf("atoi错误")
			}
			fmt.Println(totalPriceNum,curSoldNum,num,price)
			db.Debug().Model(&j).Where("juice_id = ?",j.JuiceId).Updates(map[string]string{
				"juice_sold_number"   :strconv.Itoa(int(curSoldNum + num)),
				"selling_total_price" :strconv.FormatFloat(totalPriceNum + float64(num) * price,'f',2,64),
				"cur_evaluate":order.CurEvaluate,
				"last_ordering_time":order.OrderingTime,
			})
			if(order.CurEvaluate == "好评"){
				curEvalNum, _ := strconv.ParseInt(j.GoodEvaluateNum,10,64)
				fmt.Println(curEvalNum)
				db.Model(&j).Where("juice_id = ?",j.JuiceId).Update("good_evaluate_num" ,strconv.Itoa(int(curEvalNum + 1)))
				fmt.Println("好评数据已更新")
			}
			fmt.Println("饮品订单信息已更新")

		}
	}
	return nil
}
func DeleteJuiceDataWithOrder(order *Order)error{
	db:=common.GetDB()
	db.AutoMigrate(&Juice{})
	for k,v := range parseJuiceInfo(order){
		j:=new(Juice)
		dbq:=db.First(&j, "juice_name = ?", k)
		fmt.Println("查询到此饮品数据：",j)

		if dbq.RowsAffected<1{//没查到
			return fmt.Errorf("无此饮料数据")
		} else if(dbq.RowsAffected == 1){ //有饮品信息，更新饮品信息

			totalPriceNum, err1 := strconv.ParseFloat(j.SellingTotalPrice,64)//当前销售额
			curSoldNum, err2 := strconv.ParseInt(j.JuiceSoldNumber,10,64)//当前已购买数量
			num,err3:=strconv.ParseInt(v,10,64)//饮品数量
			price, err4 := strconv.ParseFloat(j.Price,64)//饮品售价
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				fmt.Println("atoi错误",err1,err2,err3,err4)
				return fmt.Errorf("atoi错误")
			}
			fmt.Println(totalPriceNum,curSoldNum,num,price)
			db.Debug().Model(&j).Updates(map[string]string{
				"juice_sold_number"   :strconv.Itoa(int(curSoldNum - num)),
				"selling_total_price" :strconv.FormatFloat(totalPriceNum - float64(num) * price,'f',2,64),
				"cur_evaluate":"--",
				"last_ordering_time":"--",
			})
			if(order.CurEvaluate == "好评"){
				curEvalNum, _ := strconv.ParseInt(j.GoodEvaluateNum,10,64)
				fmt.Println(curEvalNum)
				db.Model(&j).Update("good_evaluate_num" ,strconv.Itoa(int(curEvalNum - 1)))
				fmt.Println("好评数据已更新")
			}
			fmt.Println("饮品订单信息已更新")

		}
	}
	return nil
}

func RefreshJucieInfo() ([]Juice,error){
	db:=common.GetDB()
	db.AutoMigrate(&Juice{})
	juices := []Juice{}
	dbs:= db.Find(&juices)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,orders)

	return juices,nil
}
func GetInexpenseOfJuice(juiceNumMapper map[string]map[string]string)((map[string]map[string]string),(map[string]map[string]string) ){
	db:=common.GetDB()
	db.AutoMigrate(&Juice{})
	juices := []Juice{}
	dbs:= db.Find(&juices)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		//return nil,err
	}

	monthTOjuicePriceMapper := make(map[string]map[string]string,0)
	for month,juiceNum := range juiceNumMapper{
		_,ok := monthTOjuicePriceMapper[month]
		if !ok{
			juiceNumMapper := make(map[string]string,0)
			juiceNumMapper["珍珠奶茶"] = juiceNum["珍珠奶茶"]
			juiceNumMapper["经典奶茶"] = juiceNum["经典奶茶"]
			juiceNumMapper["经典青汁"] = juiceNum["经典青汁"]
			juiceNumMapper["经典果茶"] = juiceNum["经典果茶"]
			monthTOjuicePriceMapper[month] = juiceNumMapper
		}
		for name,nums:= range monthTOjuicePriceMapper[month]{
			for _,juice :=range juices {
				if name == juice.JuiceName{
					nums,err1 := strconv.ParseFloat(nums,64)
					price,err2 := strconv.ParseFloat(juice.Price,64)
					//fmt.Println(nums,price)
					monthTOjuicePriceMapper[month][name] =
						strconv.FormatFloat(nums * price,'f',2,64)
					if err1 != nil || err2 != nil {
						fmt.Println("atoi错误",err1,err2)
						//return nil,nil,fmt.Errorf("atoi错误")
					}
				}
			}

		}
	}

	monthTOjuiceCostMapper := make(map[string]map[string]string,0)
	for month,juiceNum := range juiceNumMapper{
		_,ok := monthTOjuiceCostMapper[month]
		if !ok{
			juiceNumMapper := make(map[string]string,0)
			juiceNumMapper["珍珠奶茶"] = juiceNum["珍珠奶茶"]
			juiceNumMapper["经典奶茶"] = juiceNum["经典奶茶"]
			juiceNumMapper["经典青汁"] = juiceNum["经典青汁"]
			juiceNumMapper["经典果茶"] = juiceNum["经典果茶"]
			monthTOjuiceCostMapper[month] = juiceNumMapper
		}
		for name,nums:= range monthTOjuiceCostMapper[month]{
			for _,juice :=range juices {
				if name == juice.JuiceName{
					nums,err1 := strconv.ParseFloat(nums,64)
					cost,err2 := strconv.ParseFloat(juice.Cost,64)
					//fmt.Println(nums,cost)

					monthTOjuiceCostMapper[month][name] =
						strconv.FormatFloat(nums * cost,'f',2,64)
					if err1 != nil || err2 != nil {
						fmt.Println("atoi错误",err1,err2)
						//return nil,nil,fmt.Errorf("atoi错误")
					}
				}
			}

		}
	}
	return monthTOjuicePriceMapper,monthTOjuiceCostMapper
}