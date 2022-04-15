package controller

import (
	"github.com/gin-gonic/gin"
	"milkTea/common"
	"milkTea/service"
)

func refreshOwnerFunc(ctx *gin.Context){
	owner:= service.Owner{}
	//获取参数
	err:=ctx.ShouldBindJSON(&owner)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}

	owners ,err:= service.RefreshOwnerInfo(&owner)
	if err != nil {
		common.Fail(ctx,"刷新数据失败 "+err.Error(),nil)
		return
	}else {
		common.SuccessDatas(ctx,"刷新成功",owners)
	}
}

func queryOwnerFunc(ctx *gin.Context){

}

func modifyOwnerFunc(ctx *gin.Context){

}
