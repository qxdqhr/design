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

func Response(ctx *gin.Context,httpStatus int,code int,msg string,data gin.H){
	ctx.JSON(httpStatus,gin.H{"code":code, "msg":msg, "data":data,})
}
func Success(ctx *gin.Context,msg string,data gin.H){
	Response(ctx,http.StatusOK,200,msg,data)
}
func Fail(ctx *gin.Context,msg string,data gin.H){
	Response(ctx,http.StatusBadRequest,400,msg,data)
}





