package service

import (
	"fmt"
	"milkTea/common"
	"time"
)

type Alert struct {
	Id string `json:"id"`
	AlertTime string `json:"alert_time"`
	AlertReason  string `json:"alert_reason"`
	AlertMethod  string `json:"alert_method"`
	AlertOwner   string `json:"alert_owner"`
	AlertExOwner string `json:"alert_ex_owner"`
	AlertReceived string `json:"alert_received"`
}
func AddAlertInfo(alert Alert)error{
	db:=common.GetDB()
	db.AutoMigrate(&Alert{})
	a:=Alert{}
	dbq := db.Where("alert_time = ?", alert.AlertTime).First(&a)
	if(dbq.RowsAffected >= 1){
		return nil
	}else if(dbq.RowsAffected == 0){
		dbc := db.Debug().Create(&alert)
		if err := dbc.Error; err!=nil || dbc.RowsAffected < 1{
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func CheckReceiveAlert()bool {
	db:=common.GetDB()
	db.AutoMigrate(&Alert{})
	a := Alert{}
	dbs := db.Where("id = ?", "A_AUTO_" + time.Now().Format("2006-01")).First(&a)
	if dbs.RowsAffected >=1{
		return false
	}
	return true
}
func RefreshAlertInfo(userid string) ([]Alert,error){
	db:=common.GetDB()
	db.AutoMigrate(&Alert{})
	alerts := []Alert{}
	dbs:= db.Debug().Where("alert_owner = ? AND alert_received = ?",userid,"未接受").Find(&alerts)
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return nil,err
	}
	//fmt.Println("all orders of ",userid,orders)

	return alerts,nil
}
func RecieveAlertInfo(alert Alert) (error){
	db:=common.GetDB()
	db.AutoMigrate(&Alert{})

	dbs:= db.Debug().Model(&Alert{}).Where("alert_owner = ? AND alert_received = ? AND alert_reason = ? AND alert_time = ?",
		alert.AlertOwner,"未接受",alert.AlertReason,alert.AlertTime).Updates(map[string]interface{}{
		"alert_received":"已接受",
	})
	if err := dbs.Error; err!=nil {
		fmt.Println(err)
		return err
	}
	//fmt.Println("all orders of ",userid,orders)

	return nil
}