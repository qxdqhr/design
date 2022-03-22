package controller

import (
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
}

func queryOrderFunc(ctx *gin.Context){

}

func deleteFunc(ctx *gin.Context){

}

func modifyOrderFunc(ctx *gin.Context){

}

func refreshFunc(ctx *gin.Context){

}
