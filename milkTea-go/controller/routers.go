package controller

import (
	"github.com/gin-gonic/gin"
	//"milkTea/common"
)


func Routers(r *gin.Engine){
	//r.Use(common.CORS())
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
		order.POST("/delete",deleteFunc)
		order.POST("/modify",modifyOrderFunc)
		order.POST("/refresh",refreshFunc)
	}

	customer:=r.Group("/customer/")
	{
		customer.POST("/queryorder",queryOrderFunc)
		customer.POST("/refresh",refreshFunc)
	}


}


