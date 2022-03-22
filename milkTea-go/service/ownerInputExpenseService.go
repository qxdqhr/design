package service

import "github.com/jinzhu/gorm"

type OwnerInputExpense struct {
	gorm.Model
	Userid    string `json:"userid" gorm:"primaryKey"`
	Month      string `json:"month"`
	TotalIncome string `json:"total_income"`
	TotalExpence string `json:"total_expence"`
	MilkTeaIncome string `json:"milk_tea_income"`
	MilkTeaExpence string `json:"milk_tea_expence"`
	FruitTeaIncome string  `json:"fruit_tea_income"`
	FruitTeaExpence string `json:"fruit_tea_expence"`
	VegetableTeaIncome string `json:"vegetable_tea_income"`
	VegetableTeaExpence string `json:"vegetable_tea_expence"`
	OtherExpence string `json:"other_expence"`
}