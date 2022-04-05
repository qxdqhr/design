package controller

import (
	"github.com/gin-gonic/gin"
	"milkTea/common"
	"milkTea/service"
)

func queryCustomerFunc(ctx *gin.Context){
	query := service.Query{}
	err:=ctx.ShouldBindJSON(&query)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}
	err = customerQueryValidate(&query)
	if err!=nil {
		common.Fail(ctx, "validate error: "+err.Error(), nil)
		return
	}

	customer:= service.Customer{}
	customers:= []service.Customer{}
	customer.UserId = query.Userid
	switch query.QueryName {
	case "顾客名":
		customer.Name = query.QueryValue
		customers,err = service.QueryCustomerName(&customer)
		break
	case "最近购买时间":
		customer.BuyingTime = query.QueryValue
		customers,err = service.QueryCustomerBuyingTime(&customer)
		break
	case "用户最近评价":
		customer.RecentEvaluate = query.QueryValue
		customers,err = service.QueryCustomerRecentEvaluate(&customer)
		break
	}
	if err != nil {
		common.Fail(ctx,"查询数据失败 "+err.Error(),nil)
		return
	}else if len(customers) == 0{
		common.SuccessDatas(ctx,"未查到数据",customers)
	}else{
		common.SuccessDatas(ctx,"查询数据成功",customers)
	}

}

func refreshCustomerFunc(ctx *gin.Context){
	customer:= service.Customer{}
	//获取参数
	err:=ctx.ShouldBindJSON(&customer)
	if err!=nil{
		common.Err("bind error"+err.Error())
		common.Fail(ctx,"bind error"+err.Error(),nil)
		return
	}

	customers,err := service.RefreshCustomerInfo(customer.UserId)

	if err != nil {
		common.Fail(ctx,"刷新数据失败 "+err.Error(),nil)
		return
	}else {
		common.SuccessDatas(ctx,"刷新成功",customers)
	}
}