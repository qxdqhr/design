package service

import (
	"fmt"
	"milkTea/common"
)

type Owner struct {
	Userid    string `json:"user_id" gorm:"column:user_id"`
	AlertTimes int `json:"alert_times"`
	RecentReason string `json:"recent_reason"`
	ExOwnerUserid string `json:"ex_owner_userid" gorm:"column:ex_owner_userid"`
}


func AddOwnerService(owner Owner)error{
	db:=common.GetDB()
	db.AutoMigrate(&Owner{})
	o := Owner{}
	dbq := db.Where("user_id = ?", owner.Userid).First(&o)
	if(dbq.RowsAffected >= 1){
		fmt.Println("商户已存在")
		return nil
	}else if(dbq.RowsAffected == 0){
		dbc := db.Debug().Create(&owner)
		if err := dbc.Error; err!=nil || dbc.RowsAffected < 1{
			fmt.Println(err)
			return err
		}
	}
	return nil
}
func RefreshOwnerInfo(owner *Owner)([]Owner,error){
	db:=common.GetDB()
	db.AutoMigrate(&Owner{})
	o := []Owner{}
	dbq := db.Debug().Where("ex_owner_userid = ?", owner.ExOwnerUserid).First(&o)
	if(dbq.RowsAffected >= 1){
		fmt.Println("商户存在")
		return o,nil
	}else if(dbq.RowsAffected == 0){
		fmt.Println("商户不存在")
		return o,nil
	}
	if err := dbq.Error; err!=nil{
		fmt.Println(err)
		return nil,err
	}
	return nil,nil
}





