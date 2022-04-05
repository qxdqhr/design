package service

import (
	"fmt"
	"milkTea/common"
)

type InputExpense struct {
	Id string `json:"id"`
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
func RefreshInexpenseInfo(userid string) ([]InputExpense,error){
	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inexpenses := []InputExpense{}
	dbs:= db.Where("userid = ?",userid).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}
func AddInexpenseInfo(inputExpense []*InputExpense) (error){
	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inputExpenses := make([]InputExpense,0)
	for _,ele := range inputExpense{
		inputExpenses = append(inputExpenses,*ele)
	}
	fmt.Println(inputExpenses)
	dbc := db.Model(&InputExpense{}).Create(&inputExpenses)
	if err := dbc.Error; err!=nil || dbc.RowsAffected < 1{
		fmt.Println(err)
		return err
	}

	return nil


}
func DeleInexpenseInfo() (error){
	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	dbd:=db.Where("1 = 1").Delete(&InputExpense{})
	if err := dbd.Error; err!=nil || dbd.RowsAffected < 1{
		fmt.Println(err)
		return err
	}


	return nil


}
func QueryInputExpenseMonth(inexpense *InputExpense) ([]InputExpense,error){

	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inexpenses := []InputExpense{}
	dbs:= db.Where("userid = ? AND month = ?",inexpense.Userid,inexpense.Month).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}
func QueryInputExpenseTotalIncome(inexpense *InputExpense) ([]InputExpense,error){

	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inexpenses := []InputExpense{}
	dbs:= db.Where("userid = ? AND total_income = ?",inexpense.Userid,inexpense.TotalIncome).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}
func QueryInputExpenseTotalExpence(inexpense *InputExpense) ([]InputExpense,error){

	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inexpenses := []InputExpense{}
	dbs:= db.Where("userid = ? AND total_expence = ?",inexpense.Userid,inexpense.TotalExpence).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}
func QueryInputExpenseMilkTeaIncome(inexpense *InputExpense) ([]InputExpense,error){

	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inexpenses := []InputExpense{}
	dbs:= db.Where("userid = ? AND milk_tea_income = ?",inexpense.Userid,inexpense.MilkTeaIncome).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}
func QueryInputExpenseMilkTeaExpence(inexpense *InputExpense) ([]InputExpense,error){

	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inexpenses := []InputExpense{}
	dbs:= db.Where("userid = ? AND milk_tea_expence = ?",inexpense.Userid,inexpense.MilkTeaExpence).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}
func QueryInputExpenseFruitTeaIncome(inexpense *InputExpense) ([]InputExpense,error){

	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inexpenses := []InputExpense{}
	dbs:= db.Where("userid = ? AND fruit_tea_income = ?",inexpense.Userid,inexpense.FruitTeaIncome).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}
func QueryInputExpenseFruitTeaExpence(inexpense *InputExpense) ([]InputExpense,error){

	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inexpenses := []InputExpense{}
	dbs:= db.Where("userid = ? AND fruit_tea_expence = ?",inexpense.Userid,inexpense.FruitTeaExpence).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}
func QueryInputExpenseVegetableTeaIncome(inexpense *InputExpense) ([]InputExpense,error){

	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inexpenses := []InputExpense{}
	dbs:= db.Where("userid = ? AND vegetable_tea_income = ?",inexpense.Userid,inexpense.VegetableTeaIncome).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}
func QueryInputExpenseVegetableTeaExpence(inexpense *InputExpense) ([]InputExpense,error){

	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inexpenses := []InputExpense{}
	dbs:= db.Where("userid = ? AND vegetable_tea_expence = ?",inexpense.Userid,inexpense.VegetableTeaExpence).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}
func QueryInputExpenseOtherExpence(inexpense *InputExpense) ([]InputExpense,error){

	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inexpenses := []InputExpense{}
	dbs:= db.Where("userid = ? AND other_expence = ?",inexpense.Userid,inexpense.OtherExpence).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}


