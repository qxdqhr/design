package service

import "github.com/jinzhu/gorm"

type Material struct {
	gorm.Model
	MaterialId               string `json:"material_id"`
	MaterialName             string `json:"material_name"`       //原料名称
	MaterialNumber           string `json:"material_number"`       //本次购入原料数量
	PerPrice                 string `json:"per_price"`      //原料单价
	MonthBuyingTime          string `json:"material_month_buying_time"`      //购入时间
	MonthTotalPrice          string `json:"material_month_total_price"`      //购入总成本
}
type MaterialJuice struct {
	gorm.Model
	MaterialId               string `json:"material_id"`
	JuiceId                  string `json:"juice_id"`
}