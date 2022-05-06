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

func deleteOwnerFunc(ctx *gin.Context){
	owner:= service.Owner{}
	//获取参数
	err:=ctx.ShouldBindJSON(&owner)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}

	err = service.DeleteOwnerInfo(&owner)
	//user
	err = service.DeleteOwnerUserInfo(owner.Userid)
	//juice
	err = service.DeleteOwnerJuiceInfo(owner.Userid)
	//inex
	err = service.DeleteOwnerInExInfo(owner.Userid)
	//material
	err = service.DeleteOwnerMaterialInfo(owner.Userid)
	//orders
	err = service.DeleteOwnerOrdersInfo(owner.Userid)

	if err != nil {
		common.Fail(ctx,"刷新数据失败 "+err.Error(),nil)
		return
	}else {
		common.SuccessDatas(ctx,"刷新成功",nil)
	}
}
