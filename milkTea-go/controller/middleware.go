package controller

import (
	"github.com/gin-gonic/gin"
	"milkTea/common"
	"milkTea/service"
	"strings"
)

func LoginAuthorize() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		//获取认证头
		tokenString:=ctx.GetHeader("Authorization")
		//判断认证格式
		if tokenString =="" || !strings.HasPrefix(tokenString,"Bearer ") {
			common.Fail(ctx,"登录校验失败",nil)
			ctx.Abort()
			return
		}
		//截取Authorization的Bearer之后的token 并解析
		tokenString=tokenString[7:]
		token, claims, err := common.AnalyzeToken(tokenString)
		if err != nil || !token.Valid{
			common.Fail(ctx,"登录校验失败",nil)
			ctx.Abort()
			return
		}
		//校验userid
		user:= service.User{Userid: claims.UserId}
		u:=service.LoginAuthorizeService(&user)
		if u==nil{
			common.Fail(ctx,"此用户不存在",nil)
			ctx.Abort()
			return
		}
		//将user存入ctx中
		ctx.Set("user",u)
		ctx.Next()

	}
}
