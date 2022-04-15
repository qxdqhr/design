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

	customers,err := service.RefreshJucieInfo(juice.UserId)

	if err != nil {
		common.Fail(ctx,"刷新数据失败 "+err.Error(),nil)
		return
	}else {
		common.SuccessDatas(ctx,"刷新成功",customers)
	}
}
func refreshExJuiceFunc (ctx *gin.Context) {
	juice:= service.Juice{}
	//获取参数
	err:=ctx.ShouldBindJSON(&juice)
	if err!=nil{
		common.Err("bind error"+err.Error())
		common.Fail(ctx,"bind error"+err.Error(),nil)
		return
	}
	//清除原来的收入数据
	err = service.ClearExownerJuiceData(juice.UserId)
	//查询当前所有子分店
	owners, err := service.GetAllSubOwner(juice.UserId)
	//更新新的榜单数据

	juices,err := service.GetExOwnerJucieInfo(owners,juice.UserId)
	//刷新新的榜单数据

	if err != nil {
		common.Fail(ctx,"刷新数据失败 "+err.Error(),nil)
		return
	}else {
		common.SuccessDatas(ctx,"刷新成功",juices)
	}

}
