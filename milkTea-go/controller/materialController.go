package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"milkTea/common"
	"milkTea/service"
)

func addMaterialFunc(ctx *gin.Context){
	material:= service.Material{}
	//获取参数
	err:=ctx.ShouldBindJSON(&material)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}
	//校验信息
	err = materialValidate(material)
	if err != nil {
		common.Fail(ctx,"校验失败 "+err.Error(),nil)
		return
	}
	err = service.AddMaterialInfo(&material)

	if err != nil {
		common.Fail(ctx,"更新数据失败 "+err.Error(),nil)
		return
	}else {
		common.Success(ctx,"",nil)
	}
}



func queryMaterialFunc(ctx *gin.Context){
	//先获取参数
	query := service.Query{}
	err:=ctx.ShouldBindJSON(&query)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}
	err = materialQueryValidate(&query)
	if err!=nil {
		common.Fail(ctx, "validate error: "+err.Error(), nil)
		return
	}

	material:= service.Material{}
	materials:= []service.Material{}
	material.UserId = query.Userid
	switch query.QueryName {
	case "原料名称":
		material.MaterialName = query.QueryValue
		materials,err = service.QueryMaterialName(&material)
		break
	case "购入数量":
		material.MaterialNumber = query.QueryValue
		materials,err = service.QueryMaterialNumber(&material)
		break
	case "原料单价":
		material.PerPrice = query.QueryValue
		materials,err = service.QueryMaterialPerPrice(&material)
		break
	case "购入时间":
		material.MonthBuyingTime = query.QueryValue
		materials,err = service.QueryMaterialBuyingTime(&material)
		break
	case "购入成本":
		material.MonthTotalPrice = query.QueryValue
		materials,err = service.QueryMaterialTotalBuyingPrice(&material)
		break
	}
	if err != nil {
		common.Fail(ctx,"查询数据失败 "+err.Error(),nil)
		return
	}else if len(materials) == 0{
		common.SuccessDatas(ctx,"未查到数据",materials)
	}else{
		common.SuccessDatas(ctx,"查询数据成功",materials)
	}

}

func deleteMaterialFunc(ctx *gin.Context){
	material:= service.Material{}
	//获取参数
	err:=ctx.ShouldBindJSON(&material)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}
	//校验信息
	err = materialValidate(material)
	if err != nil {
		common.Fail(ctx,"校验失败 "+err.Error(),nil)
		return
	}
	err = service.DeleteMaterialInfo(&material)

	if err != nil {
		common.Fail(ctx,"删除数据失败 "+err.Error(),nil)
		return
	}else {
		common.Success(ctx,"",nil)
	}
}

func modifyMaterialFunc(ctx *gin.Context){
	material:= service.Material{}
	//获取参数
	err:=ctx.ShouldBindJSON(&material)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}
	fmt.Print(material)
	//校验信息
	err = materialValidate(material)
	if err != nil {
		common.Fail(ctx,"校验失败 "+err.Error(),nil)
		return
	}
	err = service.ModifyMaterialInfo(&material)

	if err != nil {
		common.Fail(ctx,"修改数据失败 "+err.Error(),nil)
		return
	}else {
		common.Success(ctx,"",nil)
	}
}

func refreshMaterialFunc(ctx *gin.Context){
	material:= service.Material{}
	//获取参数
	err:=ctx.ShouldBindJSON(&material)
	if err!=nil{
		common.Err("bind error"+err.Error())
		common.Fail(ctx,"bind error"+err.Error(),nil)
		return
	}
	materials,err := service.RefreshMaterialInfo(material.UserId)

	if err != nil {
		common.Fail(ctx,"刷新数据失败 "+err.Error(),nil)
		return
	}else {
		common.SuccessDatas(ctx,"刷新成功",materials)
	}
}

