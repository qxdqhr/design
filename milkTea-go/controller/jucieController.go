package controller

import (
	"github.com/gin-gonic/gin"
	"milkTea/common"
	"milkTea/service"
)

func refreshJuiceFunc(ctx *gin.Context){
	juice:= service.Juice{}
	//获取参数
	err:=ctx.ShouldBindJSON(&juice)
	if err!=nil{
		common.Err("bind error"+err.Error())
		common.Fail(ctx,"bind error"+err.Error(),nil)
		return
	}

	customers,err := service.RefreshJucieInfo()

	if err != nil {
		common.Fail(ctx,"刷新数据失败 "+err.Error(),nil)
		return
	}else {
		common.SuccessDatas(ctx,"刷新成功",customers)
	}
}
