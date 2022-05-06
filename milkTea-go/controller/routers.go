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
		juice.POST("/exowner/refresh", refreshExJuiceFunc)
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
	exinexpense:=r.Group("/exinexpense/")
	{
		exinexpense.POST("/refresh", refreshExInexpenseFunc)
	}
	alert:=r.Group("/alert/")
	{
		alert.POST("/refresh", refreshAlertFunc)
		alert.POST("/query", queryAlertFunc)
		alert.POST("/receive", receiveAlertFunc)
		alert.POST("/exOwner/add", sendAlertFunc)

	}
	owner:=r.Group("/owner/")
	{
		owner.POST("/refresh", refreshOwnerFunc)
		owner.POST("/delete", deleteOwnerFunc)
	}

}


