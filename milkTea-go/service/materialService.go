package service

import (
	"fmt"
	"gorm.io/gorm"
	"milkTea/common"
	"strconv"
)

type Material struct {
	MaterialId               string `json:"material_id"`
	MaterialName             string `json:"material_name"`       //原料名称
	MaterialNumber           string `json:"material_number"`       //本次购入原料数量
	PerPrice                 string `json:"per_price"`      //原料单价
	MonthBuyingTime          string `json:"material_month_buying_time"`      //购入时间
	MonthTotalPrice          string `json:"material_month_total_price"`      //购入总成本
	UserId  				 string  `json:"user_id"`
}
type MaterialJuice struct {
	gorm.Model
	MaterialId               string `json:"material_id"`
	JuiceId                  string `json:"juice_id"`
}
func GetInexpenseOfMaterial(userid string)(map[string]string,error){
	db:=common.GetDB()
	db.AutoMigrate(&Material{})
	materials := []Material{}
	dbs:= db.Where("user_id = ?",userid).Find(&materials)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
	}
	monthCost:=make(map[string]string,0)
	for _,material := range materials{
		//计算月份
		month := material.MonthBuyingTime[0:7]
		_,ok := monthCost[month]
		if !ok{
			monthCost[month] = "0"
		}
		orginNum,err1 := strconv.ParseFloat(monthCost[month],64)
		costNum,err2 := strconv.ParseFloat(material.MonthTotalPrice,64)
		if err1 != nil || err2 != nil {
			fmt.Println("atoi错误",err1,err2)
			return nil,fmt.Errorf("atoi错误")
		}
		monthCost[month] = strconv.FormatFloat(costNum + orginNum,'f',2,64)
	}
	return monthCost,nil
}

func AddMaterialInfo(material *Material)(error){
	db:=common.GetDB()
	db.AutoMigrate(&Material{})

	dbc := db.Debug().Create(material)
	if err := dbc.Error; err!=nil || dbc.RowsAffected < 1{
		fmt.Println(err)
		return err
	}
	return nil
}
func QueryMaterialName(material *Material) ([]Material,error) {
	db:=common.GetDB()
	db.AutoMigrate(&Material{})
	materials := []Material{}
	dbs:= db.Where("user_id = ? AND material_name = ?",material.UserId,material.MaterialName).Find(&materials)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all materials of ",userid,materials)

	return materials,nil
}

func QueryMaterialNumber(material *Material) ([]Material,error)  {
	db:=common.GetDB()
	db.AutoMigrate(&Material{})
	materials := []Material{}
	dbs:= db.Where("user_id = ? AND material_number = ?",material.UserId,material.MaterialNumber).Find(&materials)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all materials of ",userid,materials)

	return materials,nil
}

func QueryMaterialBuyingTime(material *Material) ([]Material,error)  {
	db:=common.GetDB()
	db.AutoMigrate(&Material{})
	materials := []Material{}
	dbs:= db.Where("user_id = ? AND material_month_buying_time = ?",material.UserId,material.MonthBuyingTime).Find(&materials)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all materials of ",userid,materials)

	return materials,nil
}


func QueryMaterialTotalBuyingPrice(material *Material) ([]Material,error)  {
	db:=common.GetDB()
	db.AutoMigrate(&Material{})
	materials := []Material{}
	dbs:= db.Where("user_id = ? AND material_month_total_price = ?",material.UserId,material.MonthTotalPrice).Find(&materials)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all materials of ",userid,materials)

	return materials,nil
}


func QueryMaterialPerPrice(material *Material) ([]Material,error)  {
	db:=common.GetDB()
	db.AutoMigrate(&Material{})
	materials := []Material{}
	dbs:= db.Where("user_id = ? AND per_price = ?",material.UserId,material.PerPrice).Find(&materials)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all materials of ",userid,materials)

	return materials,nil
}

func ModifyMaterialInfo(material *Material) (error){
	db:=common.GetDB()
	db.AutoMigrate(&Material{})
	o := Material{}
	dbs:= db.Where("material_id = ? ",material.MaterialId).Find(&o)

	if err := dbs.Error; err!=nil || dbs.RowsAffected == 0 {
		fmt.Println(err)
		return err
	}
	//查到了，更新数据
	dbu := db.Debug().Where("material_id = ?", material.MaterialId).Updates(
		Material{
		MaterialId:      material.MaterialId,
		MaterialName:    material.MaterialName,
		MaterialNumber:  material.MaterialNumber,
		PerPrice:        material.PerPrice,
		MonthBuyingTime: material.MonthBuyingTime,
		MonthTotalPrice: material.MonthTotalPrice,
		UserId:          material.UserId,
	})
	if err := dbu.Error; err!=nil || dbu.RowsAffected <= 0 {
		fmt.Println(err)
		return err
	}
	return nil
}
func DeleteMaterialInfo(material *Material) (error){
	db:=common.GetDB()
	db.AutoMigrate(&Material{})
	m := Material{}
	dbs:= db.Where("material_id = ? ",material.MaterialId).Find(&m)

	if err := dbs.Error; err!=nil || dbs.RowsAffected == 0 {
		fmt.Println(err)
		return err
	}
	//查到了，更新数据
	dbu := db.Where("material_id = ?", material.MaterialId).Delete(&Material{})
	if err := dbu.Error; err!=nil || dbu.RowsAffected <= 0 {
		fmt.Println(err)
		return err
	}
	return nil
}
func RefreshMaterialInfo(userid string) ([]Material,error){
	db:=common.GetDB()
	db.AutoMigrate(&Material{})
	materials := []Material{}
	dbs:= db.Where("user_id = ?",userid).Find(&materials)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,orders)

	return materials,nil
}

func DeleteOwnerMaterialInfo(userId string) (error){
	db:=common.GetDB()
	db.AutoMigrate(&Material{})

	//查到了，更新数据
	dbu := db.Where("user_id = ?",userId).Delete(&Material{})
	if err := dbu.Error; err!=nil || dbu.RowsAffected <= 0 {
		fmt.Println(err)
		return err
	}
	return nil
}