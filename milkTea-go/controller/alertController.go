package controller

import (
	"github.com/gin-gonic/gin"
	"milkTea/common"
	"milkTea/service"
)

func refreshAlertFunc(ctx *gin.Context){

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

func queryAlertFunc(ctx *gin.Context){
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

func modifyAlertFunc(ctx *gin.Context){
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


