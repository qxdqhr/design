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
	fmt.Println(juices)
	for  j:=0;j< len(juices);j++{
		juice := strings.Split(juices[j],":")
		fmt.Println(juice)
		juiceNumMap[juice[0]] = juice[1]
	}
	fmt.Println(juiceNumMap)
	return juiceNumMap
}
func UpdateJuiceWithOrder(order *Order)error{
	db:=common.GetDB()
	db=db.AutoMigrate(&Juice{})
	for k,v := range parseJuiceInfo(order){
		j:=new(Juice)
		dbq:=db.First(&j, "juice_name = ?", k)
		fmt.Println("查询到此饮品数据：",j)
		defer db.Close()

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
				"juice_sold_number"   :strconv.Itoa(int(curSoldNum + num)),
				"selling_total_price" :strconv.FormatFloat(totalPriceNum + float64(num) * price,'f',2,64),
				"cur_evaluate":order.CurEvaluate,
				"last_ordering_time":order.OrderingTime,
			})
			if(order.CurEvaluate == "好评"){
				curEvalNum, _ := strconv.ParseInt(j.GoodEvaluateNum,10,64)
				fmt.Println(curEvalNum)
				db.Model(&j).Update("good_evaluate_num" ,strconv.Itoa(int(curEvalNum + 1)))
				fmt.Println("好评数据已更新")
			}
			fmt.Println("饮品订单信息已更新")

		}
	}
	return nil
}
