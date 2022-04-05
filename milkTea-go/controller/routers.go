package controller

import (
	"github.com/gin-gonic/gin"
)
func Routers(r *gin.Engine){
	login:=r.Group("/admin/")
	{
		login.POST("/register",registerFunc)
		login.POST("/login",loginFunc)
		login.POST("/authInfo",LoginAuthorize(),authInfoFunc)//获取用户信息接口
	}
	order:=r.Group("/order/")
	{
		order.POST("/add",addOrderFunc)
		order.POST("/query",queryOrderFunc)
		order.POST("/delete", deleteOrderFunc)
		order.POST("/modify",modifyOrderFunc)
		order.POST("/refresh", refreshOrderFunc)
	}
	customer:=r.Group("/customer/")
	{
		customer.POST("/query",queryCustomerFunc)
		customer.POST("/refresh", refreshCustomerFunc)
	}
	juice:=r.Group("/juice/")
	{
		juice.POST("/refresh", refreshJuiceFunc)
	}
	material:=r.Group("/material/")
	{
		material.POST("/add",addMaterialFunc)
		material.POST("/query",queryMaterialFunc)
		material.POST("/delete", deleteMaterialFunc)
		material.POST("/modify",modifyMaterialFunc)
		material.POST("/refresh", refreshMaterialFunc)

	}

	inexpense:=r.Group("/inexpense/")
	{
		inexpense.POST("/add", addInexpenseFunc)
		inexpense.POST("/refresh", refreshInexpenseFunc)
		inexpense.POST("/query",queryInexpenseFunc)
	}
	alert:=r.Group("/alert/")
	{
		alert.POST("/refresh", refreshAlertFunc)
		alert.POST("/query", queryAlertFunc)
		alert.POST("/detail", modifyAlertFunc)
	}
	owner:=r.Group("/owner/")
	{
		owner.POST("/refresh", refreshOwnerFunc)
		owner.POST("/query", queryOwnerFunc)
		owner.POST("/detail", modifyOwnerFunc)
	}

}


