package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"milkTea/common"
	"milkTea/service"
	"strconv"
)

func addInexpenseFunc(ctx *gin.Context){
	service.DeleInexpenseInfo()
	inexpense:= service.InputExpense{}
	//获取参数
	err:=ctx.ShouldBindJSON(&inexpense)
	if err!=nil{
		common.Err("bind error"+err.Error())
		common.Fail(ctx,"bind error"+err.Error(),nil)
		return
	}
	fmt.Println(inexpense)
	monthPriceMapper, juiceNumMapper, err := service.GetInexpenseOfOrder(inexpense.Userid) //month，number，price
	fmt.Println("a",monthPriceMapper)
	//fmt.Println("b",juiceNumMapper)

	monthjuicePriceMapper ,monthjuiceCostMapper := service.GetInexpenseOfJuice(juiceNumMapper)

	fmt.Println("c",monthjuicePriceMapper)
	fmt.Println("d",monthjuiceCostMapper)
	monthCost,err:= service.GetInexpenseOfMaterial(inexpense.Userid)
	fmt.Println("e",monthCost)

    inexpenses := make([]*service.InputExpense,0)


    for m,p:=range monthPriceMapper{
		inex := service.InputExpense{
			Id:                  "IE_"+inexpense.Userid+m,
			Userid:              inexpense.Userid,
			Month:               m,
			TotalIncome:         p,
			TotalExpence:        "0.00",
			MilkTeaIncome:       "0.00",
			MilkTeaExpence:      "0.00",
			FruitTeaIncome:      "0.00",
			FruitTeaExpence:     "0.00",
			VegetableTeaIncome:  "0.00",
			VegetableTeaExpence: "0.00",
			OtherExpence:        "0.00",
		}
		inexpenses = append(inexpenses, &inex)
	}

	flag := false
	for m,nameTOprice:=range monthjuicePriceMapper{
		for _,ie := range inexpenses{
			if(ie.Month == m){
				flag = true
				for name,price := range nameTOprice{
					if(name == "经典奶茶"||name == "珍珠奶茶"){
						newer,err1 := strconv.ParseFloat(price,64)
						origin,err2 := strconv.ParseFloat(ie.MilkTeaIncome,64)
						if err1 != nil || err2 != nil {
							fmt.Println("atoi错误2",err1,err2)
							//return nil,nil,fmt.Errorf("atoi错误2")
						}
						ie.MilkTeaIncome = strconv.FormatFloat(newer+origin,'f',2,64)
						//fmt.Println("111",ie)
					}
					if(name == "经典果茶"){
						newer,err1 := strconv.ParseFloat(price,64)
						origin,err2 := strconv.ParseFloat(ie.FruitTeaIncome,64)
						if err1 != nil || err2 != nil {
							fmt.Println("atoi错误2",err1,err2)
							//return nil,nil,fmt.Errorf("atoi错误2")
						}
						ie.FruitTeaIncome = strconv.FormatFloat(newer+origin,'f',2,64)
					}
					if(name == "经典青汁"){
						newer,err1 := strconv.ParseFloat(price,64)
						origin,err2 := strconv.ParseFloat(ie.VegetableTeaIncome,64)
						if err1 != nil || err2 != nil {
							fmt.Println("atoi错误2",err1,err2)
							//return nil,nil,fmt.Errorf("atoi错误2")
						}
						ie.VegetableTeaIncome = strconv.FormatFloat(newer+origin,'f',2,64)
					}
				}
			}
		}
		if flag == false{
			inex := service.InputExpense{
				Id:                  "IE_"+inexpense.Userid+m,
				Userid:              inexpense.Userid,
				Month:               m,
				TotalIncome:         "0.00",
				TotalExpence:        "0.00",
				MilkTeaIncome:       "0.00",
				MilkTeaExpence:      "0.00",
				FruitTeaIncome:      "0.00",
				FruitTeaExpence:     "0.00",
				VegetableTeaIncome:  "0.00",
				VegetableTeaExpence: "0.00",
				OtherExpence:        "0.00",
			}
			for name,price := range nameTOprice{
				if(name == "经典奶茶"||name == "珍珠奶茶"){
					newer,err1 := strconv.ParseFloat(price,64)
					origin,err2 := strconv.ParseFloat(inex.MilkTeaIncome,64)
					if err1 != nil || err2 != nil {
						fmt.Println("atoi错误2",err1,err2)
						//return nil,nil,fmt.Errorf("atoi错误2")
					}
					inex.MilkTeaIncome = strconv.FormatFloat(newer+origin,'f',2,64)
				}
				if(name == "经典果茶"){
					newer,err1 := strconv.ParseFloat(price,64)
					origin,err2 := strconv.ParseFloat(inex.FruitTeaIncome,64)
					if err1 != nil || err2 != nil {
						fmt.Println("atoi错误2",err1,err2)
						//return nil,nil,fmt.Errorf("atoi错误2")
					}
					inex.FruitTeaIncome = strconv.FormatFloat(newer+origin,'f',2,64)
				}
				if(name == "经典青汁"){
					newer,err1 := strconv.ParseFloat(price,64)
					origin,err2 := strconv.ParseFloat(inex.VegetableTeaIncome,64)
					if err1 != nil || err2 != nil {
						fmt.Println("atoi错误2",err1,err2)
						//return nil,nil,fmt.Errorf("atoi错误2")
					}
					inex.VegetableTeaIncome = strconv.FormatFloat(newer+origin,'f',2,64)
				}
			}
			inexpenses = append(inexpenses, &inex)

		}
	}
	flag = false
	for m,nameTOcost:=range monthjuiceCostMapper{
		for _,ie := range inexpenses{
			if(ie.Month == m){
				flag = true
				for name,cost := range nameTOcost{
					if(name == "经典奶茶"||name == "珍珠奶茶"){
						newer,err1 := strconv.ParseFloat(cost,64)
						origin,err2 := strconv.ParseFloat(ie.MilkTeaExpence,64)
						if err1 != nil || err2 != nil {
							fmt.Println("atoi错误2",err1,err2)
							//return nil,nil,fmt.Errorf("atoi错误2")
						}
						ie.MilkTeaExpence = strconv.FormatFloat(newer+origin,'f',2,64)
					}
					if(name == "经典果茶"){
						newer,err1 := strconv.ParseFloat(cost,64)
						origin,err2 := strconv.ParseFloat(ie.FruitTeaExpence,64)
						if err1 != nil || err2 != nil {
							fmt.Println("atoi错误2",err1,err2)
							//return nil,nil,fmt.Errorf("atoi错误2")
						}
						ie.FruitTeaExpence = strconv.FormatFloat(newer+origin,'f',2,64)
					}
					if(name == "经典青汁"){
						newer,err1 := strconv.ParseFloat(cost,64)
						origin,err2 := strconv.ParseFloat(ie.VegetableTeaExpence,64)
						if err1 != nil || err2 != nil {
							fmt.Println("atoi错误2",err1,err2)
							//return nil,nil,fmt.Errorf("atoi错误2")
						}
						ie.VegetableTeaExpence = strconv.FormatFloat(newer+origin,'f',2,64)
					}
				}
			}
		}
		if flag == false{
			inex := service.InputExpense{
				Id:                  "IE_"+inexpense.Userid+m,
				Userid:              inexpense.Userid,
				Month:               m,
				TotalIncome:         "0.00",
				TotalExpence:        "0.00",
				MilkTeaIncome:       "0.00",
				MilkTeaExpence:      "0.00",
				FruitTeaIncome:      "0.00",
				FruitTeaExpence:     "0.00",
				VegetableTeaIncome:  "0.00",
				VegetableTeaExpence: "0.00",
				OtherExpence:        "0.00",
			}
			for name,cost := range nameTOcost{
				if(name == "经典奶茶"||name == "珍珠奶茶"){
					newer,err1 := strconv.ParseFloat(cost,64)
					origin,err2 := strconv.ParseFloat(inex.MilkTeaExpence,64)
					if err1 != nil || err2 != nil {
						fmt.Println("atoi错误2",err1,err2)
						//return nil,nil,fmt.Errorf("atoi错误2")
					}
					inex.MilkTeaExpence = strconv.FormatFloat(newer+origin,'f',2,64)
				}
				if(name == "经典果茶"){
					newer,err1 := strconv.ParseFloat(cost,64)
					origin,err2 := strconv.ParseFloat(inex.FruitTeaExpence,64)
					if err1 != nil || err2 != nil {
						fmt.Println("atoi错误2",err1,err2)
						//return nil,nil,fmt.Errorf("atoi错误2")
					}
					inex.FruitTeaExpence = strconv.FormatFloat(newer+origin,'f',2,64)
				}
				if(name == "经典青汁"){
					newer,err1 := strconv.ParseFloat(cost,64)
					origin,err2 := strconv.ParseFloat(inex.VegetableTeaExpence,64)
					if err1 != nil || err2 != nil {
						fmt.Println("atoi错误2",err1,err2)
						//return nil,nil,fmt.Errorf("atoi错误2")
					}
					inex.VegetableTeaExpence = strconv.FormatFloat(newer+origin,'f',2,64)
				}
			}
			inexpenses = append(inexpenses, &inex)

		}
	}
	flag = false
	for m,otherCost:=range monthCost{
		for _,ie := range inexpenses{
			if(ie.Month == m){
				flag = true
				ie.OtherExpence = otherCost
			}
		}
		if flag == false{
			inex := service.InputExpense{
				Id:                  "IE_" + inexpense.Userid + m,
				Userid:              inexpense.Userid,
				Month:               m,
				TotalIncome:         "0.00",
				TotalExpence:        "0.00",
				MilkTeaIncome:       "0.00",
				MilkTeaExpence:      "0.00",
				FruitTeaIncome:      "0.00",
				FruitTeaExpence:     "0.00",
				VegetableTeaIncome:  "0.00",
				VegetableTeaExpence: "0.00",
				OtherExpence:        otherCost,
			}
			inexpenses = append(inexpenses, &inex)
		}
	}
	for _,val:= range inexpenses{
		fmt.Println(*val)

		total,err := strconv.ParseFloat((*val).TotalExpence,64)
		vegetable,err := strconv.ParseFloat((*val).VegetableTeaExpence,64)
		fruit,err := strconv.ParseFloat((*val).FruitTeaExpence,64)
		milk,err := strconv.ParseFloat((*val).MilkTeaExpence,64)
		others,err := strconv.ParseFloat((*val).OtherExpence,64)
		if err != nil {
			fmt.Println("atoi错误2",err)
			//return nil,nil,fmt.Errorf("atoi错误2")
		}
		(*val).TotalExpence = strconv.FormatFloat(total+vegetable+fruit+milk+others,'f',2,64)
	}

	service.AddInexpenseInfo(inexpenses)
}
func refreshInexpenseFunc(ctx *gin.Context){
	inexpense:= service.InputExpense{}
	//获取参数
	err:=ctx.ShouldBindJSON(&inexpense)
	if err!=nil{
		common.Err("bind error"+err.Error())
		common.Fail(ctx,"bind error"+err.Error(),nil)
		return
	}
	Inexpenses,err := service.RefreshInexpenseInfo(inexpense.Userid)

	if err != nil {
		common.Fail(ctx,"刷新数据失败 "+err.Error(),nil)
		return
	}else {
		common.SuccessDatas(ctx,"刷新成功",Inexpenses)
	}

}
func queryInexpenseFunc(ctx *gin.Context){
	//先获取参数
	query := service.Query{}
	err:=ctx.ShouldBindJSON(&query)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}
	err = inexpenseQueryValidate(&query)
	if err!=nil {
		common.Fail(ctx, "validate error: "+err.Error(), nil)
		return
	}

	inputexpense:= service.InputExpense{}
	inputexpenses:= []service.InputExpense{}
	inputexpense.Userid = query.Userid
	switch query.QueryName {
	case "月份数":
		inputexpense.Month = query.QueryValue
		inputexpenses,err = service.QueryInputExpenseMonth(&inputexpense)
		break
	case "总收入":
		inputexpense.TotalIncome = query.QueryValue
		inputexpenses,err = service.QueryInputExpenseTotalIncome(&inputexpense)
		break
	case "总支出":
		inputexpense.TotalExpence = query.QueryValue
		inputexpenses,err = service.QueryInputExpenseTotalExpence(&inputexpense)
		break
	case "奶茶收入":
		inputexpense.MilkTeaIncome = query.QueryValue
		inputexpenses,err = service.QueryInputExpenseMilkTeaIncome(&inputexpense)
		break
	case "奶茶支出":
		inputexpense.MilkTeaExpence = query.QueryValue
		inputexpenses,err = service.QueryInputExpenseMilkTeaExpence(&inputexpense)
		break
	case "果茶收入":
		inputexpense.FruitTeaIncome = query.QueryValue
		inputexpenses,err = service.QueryInputExpenseFruitTeaIncome(&inputexpense)
		break
	case "果茶支出":
		inputexpense.FruitTeaExpence = query.QueryValue
		inputexpenses,err = service.QueryInputExpenseFruitTeaExpence(&inputexpense)
		break
	case "青汁收入":
		inputexpense.VegetableTeaIncome = query.QueryValue
		inputexpenses,err = service.QueryInputExpenseVegetableTeaIncome(&inputexpense)
		break
	case "青汁支出":
		inputexpense.VegetableTeaExpence = query.QueryValue
		inputexpenses,err = service.QueryInputExpenseVegetableTeaExpence(&inputexpense)
		break
	case "原料支出":
		inputexpense.OtherExpence = query.QueryValue
		inputexpenses,err = service.QueryInputExpenseOtherExpence(&inputexpense)
		break
	default:
		break
	}
	if err != nil {
		common.Fail(ctx,"查询数据失败 "+err.Error(),nil)
		return
	}else if len(inputexpenses) == 0{
		common.SuccessDatas(ctx,"未查到数据",inputexpenses)
	}else{
		common.SuccessDatas(ctx,"查询数据成功",inputexpenses)
	}

}
