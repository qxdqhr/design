package service

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Name    string `json:"name"`
	CustomerId    string `json:"customerid"`
	BuyingTime string `json:"buying_time"`
	RecentEvaluate string `json:"recent_evaluate"`
}