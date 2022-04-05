package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"milkTea/common"
	"milkTea/common/config"
	"milkTea/controller"
	"milkTea/service"
)

func init(){
	config.InitConfig()
}
func MakeDB(){
	db:=common.GetDB()
	db.AutoMigrate(&service.Owner{})
	db.AutoMigrate(&service.InputExpense{})
	db.AutoMigrate(&service.Customer{})
	db.AutoMigrate(&service.Order{})
	db.AutoMigrate(&service.Juice{})
	db.AutoMigrate(&service.Material{})
	db.AutoMigrate(&service.MaterialJuice{})
	db.AutoMigrate(&service.Alert{})
	db.AutoMigrate(&service.User{})
	//j:=service.Juice{
	//	JuiceId:           "",
	//	JuiceName:         "bbb",
	//	JuiceType:         "bbb",
	//	LastOrderingTime:  "2022/3/17",
	//	Price:             "15.2",
	//	Profit:            "",
	//	Cost:              "",
	//	CurEvaluate:       "好评",
	//	JuiceSoldNumber:   "0",
	//	SellingTotalPrice: "0",
	//	GoodEvaluateNum:   "1",
	//}
	//db.Create(&j)
}
func RunDemo(){
	r:=gin.Default()
	controller.Routers(r)
	r.Run(viper.GetString("server.port"))
}

func main(){
	RunDemo()
	//MakeDB()
}
