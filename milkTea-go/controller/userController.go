package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"milkTea/common"
	"milkTea/common/dto"
	"milkTea/service"
)


func loginFunc(ctx *gin.Context)  {
	//登录输入工号以及密码
	user:= service.User{}
	resUser:= &service.User{}
	//获取参数
	err:=ctx.ShouldBindJSON(&user)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}
	//校验信息
	err = loginValidate(user)
	if err != nil {
		common.Fail(ctx,"login validate error "+err.Error(),nil)
		return
	}

	//查找数据库
	err ,resUser = service.LoginService(&user)
	fmt.Println(user)
	if err != nil {
		common.Fail(ctx,err.Error(),nil)
	}else{
		//发放token
		//token, err := common.CreateToken(user.Userid)
		//if err != nil {
		//	common.Fail(ctx,"创建 token 失败",nil)
		//	return
		//}
		common.Success(ctx,"登录成功",gin.H{
			"name":resUser.Name,
			"userId":resUser.Userid,
		})
	}
	return
}
func registerFunc(ctx *gin.Context)  {
	user:= service.User{}
	//获取参数
	err:=ctx.ShouldBindJSON(&user)
	if err!=nil{
		common.Err("bind error")
		common.Fail(ctx,"bind error",nil)
		return
	}
	//校验信息
	err = registerValidate(user)
	if err != nil {
//		common.Err(err.Error())
		common.Fail(ctx,"注册校验失败 "+err.Error(),nil)
		return
	}
	//获取工号
	user.Userid=common.GetUid(user.Role,user.Telephone)

	//密码加密
	passwordEx, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.MinCost)
	if err != nil {
//		common.Err(err.Error())
		common.Fail(ctx,"密码加密失败"+err.Error(),nil)
		return
	}
	user.Password=string(passwordEx)
	//service
	err = service.RegisterService(&user)
	//返回信息:工号,和发信息
	if err != nil {
		common.Fail(ctx,err.Error(),gin.H{"userid":user.Userid})
	}else{
		common.Success(ctx,"注册成功",gin.H{"userid":user.Userid})
	}

}
func authInfoFunc(ctx *gin.Context){
	u, exists := ctx.Get("user")
	if exists {
		common.Success(ctx,"登录用户",gin.H{"login user":dto.ToReturnLoginAuthorize(u)})
	}else{
		common.Fail(ctx,"当前没有用户登陆",nil)
	}
}
