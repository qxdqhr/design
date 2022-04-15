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
	UserId string `json:"user_id"`
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
		dbq:=db.First(&j, "juice_name = ? AND user_id = ?", k,order.UserId)
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
			juice_sold_number:=strconv.Itoa(int(curSoldNum + num))
			selling_total_price:=strconv.FormatFloat(totalPriceNum + float64(num) * price,'f',2,64)
			fmt.Println("new:",juice_sold_number,selling_total_price,order.CurEvaluate,order.OrderingTime)
			db:=common.GetDB()
			db.Debug().Model(&Juice{}).Where("juice_id = ?  AND user_id = ?",j.JuiceId,order.UserId).Updates(map[string]interface{}{
				"juice_sold_number"   :juice_sold_number,
				"selling_total_price" :selling_total_price,
				"cur_evaluate":order.CurEvaluate,
				"last_ordering_time":order.OrderingTime,
			})
			if(order.CurEvaluate == "好评"){
				curEvalNum, _ := strconv.ParseInt(j.GoodEvaluateNum,10,64)
				fmt.Println(curEvalNum)
				db.Model(&j).Where("juice_id = ? AND user_id = ?",j.JuiceId,order.UserId).Update("good_evaluate_num" ,strconv.Itoa(int(curEvalNum + 1)))
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
			db.Debug().Model(&j).Updates(map[string]interface{}{
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

func RefreshJucieInfo(userid string) ([]Juice,error){
	db:=common.GetDB()
	db.AutoMigrate(&Juice{})
	juices := []Juice{}
	dbs:= db.Debug().Where("user_id = ?",userid).Find(&juices)
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
func AddJuiceInfo(userid string)error{
	
	db:=common.GetDB()
	db.AutoMigrate(&Juice{})
	juices :=make([]Juice,4)
	juices[0] = Juice{
		JuiceId:           "bubble_milkeTea",
		JuiceName:         "珍珠奶茶",
		JuiceType:         "奶茶",
		LastOrderingTime:  "--",
		Price:             "11.0",
		Profit:            "2.0",
		Cost:              "9.0",
		CurEvaluate:       "--",
		JuiceSoldNumber:   "0",
		SellingTotalPrice: "0.00",
		GoodEvaluateNum:   "0",
		UserId: userid,
	}
	juices[1] = Juice{
		JuiceId:           "classic_fruitTea",
		JuiceName:         "经典果茶",
		JuiceType:         "果茶",
		LastOrderingTime:  "--",
		Price:             "9.0",
		Profit:            "2.0",
		Cost:              "7.0",
		CurEvaluate:       "--",
		JuiceSoldNumber:   "0",
		SellingTotalPrice: "0.00",
		GoodEvaluateNum:   "0",
		UserId: userid,
	}
	juices[2] = Juice{
		JuiceId:           "classic_milkTea",
		JuiceName:         "经典奶茶",
		JuiceType:         "奶茶",
		LastOrderingTime:  "--",
		Price:             "10.0",
		Profit:            "2.0",
		Cost:              "8.0",
		CurEvaluate:       "--",
		JuiceSoldNumber:   "0",
		SellingTotalPrice: "0.00",
		GoodEvaluateNum:   "0",
		UserId: userid,
	}
	juices[3] = Juice{
		JuiceId:           "classic_vegetableJuice",
		JuiceName:         "经典青汁",
		JuiceType:         "青汁",
		LastOrderingTime:  "--",
		Price:             "12.0",
		Profit:            "3.0",
		Cost:              "9.0",
		CurEvaluate:       "--",
		JuiceSoldNumber:   "0",
		SellingTotalPrice: "0.00",
		GoodEvaluateNum:   "0",
		UserId: userid,
	}
	j:=new(Juice)
	dbq:=db.First(&j, "user_id = ?", userid)
	fmt.Println("查询到此饮品数据：",j)

	if dbq.RowsAffected >= 1{
		fmt.Println("已存在饮品数据：",j)

	}else {
		dbc := db.Debug().Create(&juices)
		if err := dbc.Error; err!=nil || dbc.RowsAffected < 1{
			fmt.Println(err)
			return err
		}
	}


	return nil
	
}
func ClearExownerJuiceData(userid string)error{

	db:=common.GetDB()
	db.AutoMigrate(&Juice{})

	j:=new(Juice)
	dbq:=db.Debug().Where(&j, "user_id = ?", userid).Updates(map[string]interface{}{
		"juice_sold_number"   :"0",
		"selling_total_price" :"0",
		"cur_evaluate":"--",
		"last_ordering_time":"--",
	})
	fmt.Println("查询到此饮品数据：",j)
	if err := dbq.Error; err!=nil || dbq.RowsAffected < 1{
		fmt.Println(err)
		return err
	}
	return nil

}
func GetExOwnerJucieInfo(owners []string,userid string)([]Juice,error){
	db:=common.GetDB()
	db.AutoMigrate(&Juice{})
	res := make ([]Juice,4)
	res[0] = Juice{
		JuiceId:           "bubble_milkeTea",
		JuiceName:         "珍珠奶茶",
		JuiceType:         "奶茶",
		LastOrderingTime:  "--",
		Price:             "11.0",
		Profit:            "2.0",
		Cost:              "9.0",
		CurEvaluate:       "--",
		JuiceSoldNumber:   "0",
		SellingTotalPrice: "0.00",
		GoodEvaluateNum:   "0",
	}
	res[1] = Juice{
		JuiceId:           "classic_fruitTea",
		JuiceName:         "经典果茶",
		JuiceType:         "果茶",
		LastOrderingTime:  "--",
		Price:             "9.0",
		Profit:            "2.0",
		Cost:              "7.0",
		CurEvaluate:       "--",
		JuiceSoldNumber:   "0",
		SellingTotalPrice: "0.00",
		GoodEvaluateNum:   "0",
	}
	res[2] = Juice{
		JuiceId:           "classic_milkTea",
		JuiceName:         "经典奶茶",
		JuiceType:         "奶茶",
		LastOrderingTime:  "--",
		Price:             "10.0",
		Profit:            "2.0",
		Cost:              "8.0",
		CurEvaluate:       "--",
		JuiceSoldNumber:   "0",
		SellingTotalPrice: "0.00",
		GoodEvaluateNum:   "0",
	}
	res[3] = Juice{
		JuiceId:           "classic_vegetableJuice",
		JuiceName:         "经典青汁",
		JuiceType:         "青汁",
		LastOrderingTime:  "--",
		Price:             "12.0",
		Profit:            "3.0",
		Cost:              "9.0",
		CurEvaluate:       "--",
		JuiceSoldNumber:   "0",
		SellingTotalPrice: "0.00",
		GoodEvaluateNum:   "0",
	}
	juices:=[]Juice{}
	dbq := db.Debug().Model(&Juice{}).Where("user_id in (?)", owners).Find(&juices)
	if err := dbq.Error; err!=nil || dbq.RowsAffected < 1{
		fmt.Println(err)
		return nil,err
	}
	for _,v:= range juices{
		switch v.JuiceName{
		case "珍珠奶茶":
			onum, _ := strconv.ParseFloat(res[0].JuiceSoldNumber, 64)
			nnum, _ := strconv.ParseFloat(v.JuiceSoldNumber,64)
			res[0].JuiceSoldNumber = strconv.FormatFloat(onum+nnum,'f',2,64)

			onum, _ = strconv.ParseFloat(res[0].SellingTotalPrice, 64)
			nnum, _ = strconv.ParseFloat(v.SellingTotalPrice,64)
			res[0].SellingTotalPrice = strconv.FormatFloat(onum+nnum,'f',2,64)

			onum, _ = strconv.ParseFloat(res[0].GoodEvaluateNum, 64)
			nnum, _ = strconv.ParseFloat(v.GoodEvaluateNum,64)
			res[0].GoodEvaluateNum = strconv.FormatFloat(onum+nnum,'f',2,64)
			break
		case "经典果茶":
			onum, _ := strconv.ParseFloat(res[1].JuiceSoldNumber, 64)
			nnum, _ := strconv.ParseFloat(v.JuiceSoldNumber,64)
			res[1].JuiceSoldNumber = strconv.FormatFloat(onum+nnum,'f',2,64)

			onum, _ = strconv.ParseFloat(res[1].SellingTotalPrice, 64)
			nnum, _ = strconv.ParseFloat(v.SellingTotalPrice,64)
			res[1].SellingTotalPrice = strconv.FormatFloat(onum+nnum,'f',2,64)

			onum, _ = strconv.ParseFloat(res[1].GoodEvaluateNum, 64)
			nnum, _ = strconv.ParseFloat(v.GoodEvaluateNum,64)
			res[1].GoodEvaluateNum = strconv.FormatFloat(onum+nnum,'f',2,64)
			break
		case "经典奶茶":
			onum, _ := strconv.ParseFloat(res[2].JuiceSoldNumber, 64)
			nnum, _ := strconv.ParseFloat(v.JuiceSoldNumber,64)
			res[2].JuiceSoldNumber = strconv.FormatFloat(onum+nnum,'f',2,64)

			onum, _ = strconv.ParseFloat(res[2].SellingTotalPrice, 64)
			nnum, _ = strconv.ParseFloat(v.SellingTotalPrice,64)
			res[2].SellingTotalPrice = strconv.FormatFloat(onum+nnum,'f',2,64)

			onum, _ = strconv.ParseFloat(res[2].GoodEvaluateNum, 64)
			nnum, _ = strconv.ParseFloat(v.GoodEvaluateNum,64)
			res[2].GoodEvaluateNum = strconv.FormatFloat(onum+nnum,'f',2,64)
			break
		case "经典青汁":
			onum, _ := strconv.ParseFloat(res[3].JuiceSoldNumber, 64)
			nnum, _ := strconv.ParseFloat(v.JuiceSoldNumber,64)
			res[3].JuiceSoldNumber = strconv.FormatFloat(onum+nnum,'f',2,64)

			onum, _ = strconv.ParseFloat(res[3].SellingTotalPrice, 64)
			nnum, _ = strconv.ParseFloat(v.SellingTotalPrice,64)
			res[3].SellingTotalPrice = strconv.FormatFloat(onum+nnum,'f',2,64)

			onum, _ = strconv.ParseFloat(res[3].GoodEvaluateNum, 64)
			nnum, _ = strconv.ParseFloat(v.GoodEvaluateNum,64)
			res[3].GoodEvaluateNum = strconv.FormatFloat(onum+nnum,'f',2,64)
			break
		default:
			break
		}
	}
	dbu:=db.Debug().Model(&Juice{}).Where("user_id = ? AND juice_id = ?",userid,"bubble_milkeTea").Updates(
		map[string]interface{}{
		"juice_sold_number"   :res[0].JuiceSoldNumber,
		"selling_total_price" :res[0].SellingTotalPrice,
		"last_ordering_time":res[0].LastOrderingTime,
	})
	dbu = db.Debug().Model(&Juice{}).Where("user_id = ? AND juice_id = ?",userid,"classic_fruitTea").Updates(
		map[string]interface{}{
			"juice_sold_number"   :res[1].JuiceSoldNumber,
			"selling_total_price" :res[1].SellingTotalPrice,
			"last_ordering_time":res[1].LastOrderingTime,
		})
	dbu = db.Debug().Model(&Juice{}).Where("user_id = ? AND juice_id = ?",userid,"classic_milkTea").Updates(
		map[string]interface{}{
			"juice_sold_number"   :res[2].JuiceSoldNumber,
			"selling_total_price" :res[2].SellingTotalPrice,
			"last_ordering_time":res[2].LastOrderingTime,
		})
	dbu = db.Debug().Model(&Juice{}).Where("user_id = ? AND juice_id = ?",userid,"classic_vegetableJuice").Updates(
		map[string]interface{}{
			"juice_sold_number"   :res[3].JuiceSoldNumber,
			"selling_total_price" :res[3].SellingTotalPrice,
			"last_ordering_time":res[3].LastOrderingTime,
		})
	if err := dbu.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	return res,nil
}
