package service

import (
	"fmt"
	"milkTea/common"
	"sort"
	"strconv"
)

type InputExpense struct {
	Id string `json:"id"`
	Userid    string `json:"user_id" gorm:"column:user_id"`
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
type IeSort []InputExpense

	//sort
func (m IeSort) Len() int {
	return len(m)
}
// 实现sort.Interface接口的比较元素方法
func (m IeSort) Less(i, j int) bool {
	return m[i].Month < m[j].Month
}
// 实现sort.Interface接口的交换元素方法
func (m IeSort) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func CheckAlert(inputExpense []InputExpense)bool{
	sort.Sort(IeSort(inputExpense))
	fmt.Println(inputExpense)
	isAlert := false
	lossCount := 0
	if len(inputExpense)<3 {
		//不会告警
		return false
	}

	for _,ele :=range inputExpense{
		income,_ :=strconv.ParseFloat(ele.TotalIncome,64)
		expense,_ :=strconv.ParseFloat(ele.TotalExpence,64)
		if(income - expense <=0 ) {
			lossCount++
		}else{
			lossCount = 0
		}
		if(lossCount >= 3){
			isAlert = true
		}
	}
	return isAlert
}
func RefreshInexpenseInfo(userid string) ([]InputExpense,error){
	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inexpenses := []InputExpense{}
	dbs:= db.Where("user_id = ?",userid).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}
func AddInexpenseInfo(inputExpense []InputExpense) (error){
	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	//fmt.Println(inputExpense)
	dbc := db.Debug().Model(&InputExpense{}).Create(&inputExpense)
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
	dbs:= db.Where("user_id = ? AND month = ?",inexpense.Userid,inexpense.Month).Find(&inexpenses)
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
	dbs:= db.Where("user_id = ? AND total_income = ?",inexpense.Userid,inexpense.TotalIncome).Find(&inexpenses)
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
	dbs:= db.Where("user_id = ? AND total_expence = ?",inexpense.Userid,inexpense.TotalExpence).Find(&inexpenses)
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
	dbs:= db.Where("user_id = ? AND milk_tea_income = ?",inexpense.Userid,inexpense.MilkTeaIncome).Find(&inexpenses)
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
	dbs:= db.Where("user_id = ? AND milk_tea_expence = ?",inexpense.Userid,inexpense.MilkTeaExpence).Find(&inexpenses)
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
	dbs:= db.Where("user_id = ? AND fruit_tea_income = ?",inexpense.Userid,inexpense.FruitTeaIncome).Find(&inexpenses)
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
	dbs:= db.Where("user_id = ? AND fruit_tea_expence = ?",inexpense.Userid,inexpense.FruitTeaExpence).Find(&inexpenses)
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
	dbs:= db.Where("user_id = ? AND vegetable_tea_income = ?",inexpense.Userid,inexpense.VegetableTeaIncome).Find(&inexpenses)
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
	dbs:= db.Where("user_id = ? AND vegetable_tea_expence = ?",inexpense.Userid,inexpense.VegetableTeaExpence).Find(&inexpenses)
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
	dbs:= db.Where("user_id = ? AND other_expence = ?",inexpense.Userid,inexpense.OtherExpence).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}

func GetExOwnerInexpenseInfo(userids []string) ([]InputExpense,error){
	db:=common.GetDB()
	db.AutoMigrate(&InputExpense{})
	inexpenses := []InputExpense{}
	dbs:= db.Where("user_id = ?",userid).Find(&inexpenses)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all inexpenses of ",userid,inexpenses)

	return inexpenses,nil
}
