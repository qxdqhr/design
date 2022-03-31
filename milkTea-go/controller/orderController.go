package controller

import (
	"fmt"
	//json "encoding/json"
	"github.com/gin-gonic/gin"
	"milkTea/common"
	"milkTea/service"
)

func addOrderFunc(ctx *gin.Context){
	order:= service.Order{}
	//获取参数
	err:=ctx.ShouldBindJSON(&order)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}
	//校验信息
	err = orderValidate(order)
	if err != nil {
		common.Fail(ctx,"校验失败 "+err.Error(),nil)
		return
	}
	err = service.AddOrderInfo(&order)

	c := service.Customer{
		Name:           order.CustomerId,
		CustomerId:     common.GetCustomerid(order.CustomerId),
		BuyingTime:     order.OrderingTime,
		RecentEvaluate: order.CurEvaluate,
	}

	err = service.UpdateCustomerInfo(&c)
	//向 juice 表中写
	err = service.UpdateJuiceWithOrder(&order)

	if err != nil {
		common.Fail(ctx,"更新数据失败 "+err.Error(),nil)
		return
	}else {
		common.Success(ctx,"",nil)
	}
}



func queryOrderFunc(ctx *gin.Context){
	//先获取参数
	query := service.Query{}
	err:=ctx.ShouldBindJSON(&query)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}
	err = orderQueryValidate(&query)
	if err!=nil {
		common.Fail(ctx, "validate error: "+err.Error(), nil)
		return
	}

	order:= service.Order{}
	orders:= []service.Order{}
	order.UserId = query.Userid
	switch query.QueryName {
	case "顾客名":
		order.CustomerId = query.QueryValue
		orders,err = service.QueryOrderCustomerName(&order)
		break
	case "购买饮品":
		order.Buyingjuice = query.QueryValue
		orders,err = service.QueryOrderBuyingjuice(&order)
		break
	case "下单时间":
		order.OrderingTime = query.QueryValue
		orders,err = service.QueryOrderOrderingTime(&order)
		break
	case "本单饮品数量":
		order.JuiceNumber = query.QueryValue
		orders,err = service.QueryOrderJuiceNumber(&order)
		break
	case "总售价":
		order.TotalSellingPrice = query.QueryValue
		orders,err = service.QueryOrderTotalSellingPrice(&order)
		break
	case "本单评价":
		order.CurEvaluate = query.QueryValue
		orders,err = service.QueryOrderCurEvaluate(&order)
		break
	}
	if err != nil {
		common.Fail(ctx,"查询数据失败 "+err.Error(),nil)
		return
	}else if len(orders) == 0{
		common.SuccessDatas(ctx,"未查到数据",orders)
	}else{
		common.SuccessDatas(ctx,"查询数据成功",orders)
	}

}

func deleteFunc(ctx *gin.Context){
	order:= service.Order{}
	//获取参数
	err:=ctx.ShouldBindJSON(&order)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}
	//校验信息
	err = orderValidate(order)
	if err != nil {
		common.Fail(ctx,"校验失败 "+err.Error(),nil)
		return
	}
	err = service.DeleteOrderInfo(&order)

	c := service.Customer{
		Name:           order.CustomerId,
		CustomerId:     common.GetCustomerid(order.CustomerId),
		BuyingTime:     order.OrderingTime,
		RecentEvaluate: order.CurEvaluate,
	}

	err = service.UpdateCustomerInfo(&c)
	//向 juice 表中写
	err = service.UpdateJuiceWithOrder(&order)

	if err != nil {
		common.Fail(ctx,"更新数据失败 "+err.Error(),nil)
		return
	}else {
		common.Success(ctx,"",nil)
	}
}

func modifyOrderFunc(ctx *gin.Context){
	order:= service.Order{}
	//获取参数
	err:=ctx.ShouldBindJSON(&order)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}
	fmt.Print(order)
	//校验信息
	err = orderValidate(order)
	if err != nil {
		common.Fail(ctx,"校验失败 "+err.Error(),nil)
		return
	}
	err = service.ModifyOrderInfo(&order)

	c := service.Customer{
		Name:           order.CustomerId,
		CustomerId:     common.GetCustomerid(order.CustomerId),
		BuyingTime:     order.OrderingTime,
		RecentEvaluate: order.CurEvaluate,
	}

	err = service.UpdateCustomerInfo(&c)
	//向 juice 表中写
	err = service.UpdateJuiceWithOrder(&order)

	if err != nil {
		common.Fail(ctx,"修改数据失败 "+err.Error(),nil)
		return
	}else {
		common.Success(ctx,"",nil)
	}
}

func refreshFunc(ctx *gin.Context){
	order:= service.Order{}
	//获取参数
	err:=ctx.ShouldBindJSON(&order)
	if err!=nil{
		common.Err("bind error"+err.Error())
		common.Fail(ctx,"bind error"+err.Error(),nil)
		return
	}
	orders,err := service.RefreshOrderInfo(order.UserId)

	if err != nil {
		common.Fail(ctx,"刷新数据失败 "+err.Error(),nil)
		return
	}else {
		common.SuccessDatas(ctx,"刷新成功",orders)
	}
}
