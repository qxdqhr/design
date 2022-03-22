package service

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	CustomerId    string `json:"customerid"`
	OrderingTime string `json:"orderingtime"`   //下单时间
	JuiceNumber string `json:"juicenumber"`     //本单饮品数量
	TotalSellingPrice string `json:"totalsellingprice"`//总售价
	CurEvaluate string  `json:"curevaluate"`    //本单评价
	Buyingjuice string `json:"buyingjuice"`
}


