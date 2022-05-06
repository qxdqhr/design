package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"milkTea/common"
	"milkTea/service"
)

func refreshAlertFunc(ctx *gin.Context){

	alert:= service.Alert{}
	//获取参数
	err:=ctx.ShouldBindJSON(&alert)
	if err!=nil{
		common.Err("bind error"+err.Error())
		common.Fail(ctx,"bind error"+err.Error(),nil)
		return
	}

	alertInfo,err := service.RefreshAlertInfo(alert.AlertOwner)
	//if len(alertInfo){
	//
	//}
	if err != nil {
		common.Fail(ctx,"刷新数据失败 "+err.Error(),nil)
		return
	}else {
		common.SuccessDatas(ctx,"刷新成功",alertInfo)
	}
}
func receiveAlertFunc(ctx *gin.Context){
	alert:= service.Alert{}
	//获取参数
	err:=ctx.ShouldBindJSON(&alert)
	if err!=nil{
		common.Err("bind error"+err.Error())
		common.Fail(ctx,"bind error"+err.Error(),nil)
		return
	}
	fmt.Println("aaa")
	err = service.RecieveAlertInfo(alert)
	if err != nil {
		common.Fail(ctx,"确认告警失败 "+err.Error(),nil)
		return
	}else {
		common.SuccessDatas(ctx,"确认告警成功",nil)
	}
}
func queryAlertFunc(ctx *gin.Context){

}

func sendAlertFunc(ctx *gin.Context){
	alert:= service.Alert{}
	//获取参数
	err:=ctx.ShouldBindJSON(&alert)
	if err!=nil{
		common.Err("bind error"+err.Error())
		common.Fail(ctx,"bind error"+err.Error(),nil)
		return
	}

	res,err := service.AddAlertInfo(alert)
	if err != nil {
		common.Fail(ctx,"确认告警失败 "+err.Error(),nil)
		return
	}else {
		common.SuccessDatas(ctx,"确认告警成功" + res,nil)
	}
}


