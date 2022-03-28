package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetUid(role,telephone string)(string){
	timeString :=time.Now().Format("060102")
	switch role{
	case "Owner":
		return telephone[7:]+"A"+timeString
	case "ExOwner":
		return telephone[7:]+"B"+timeString
	}
	return ""
}

func GetCustomerid(username string)(string){
	return "C_" + username
}

func Response(ctx *gin.Context,httpStatus int,code int,msg string,data interface{}) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
func ResponseDatas(ctx *gin.Context,httpStatus int,code int,msg string,datas interface{}){
	ctx.IndentedJSON(httpStatus,gin.H{
		"code": code,
		"msg":  msg,
		"data": datas,
	})
}
func SuccessDatas(ctx *gin.Context,msg string,datas interface{}){
	ResponseDatas(ctx,http.StatusOK,200,msg,datas)
}
func Success(ctx *gin.Context,msg string,data interface{}){
	Response(ctx,http.StatusOK,200,msg,data)
}
func Fail(ctx *gin.Context,msg string,data interface{}){
	Response(ctx,http.StatusBadRequest,400,msg,data)
}





