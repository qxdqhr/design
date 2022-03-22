package service

import "github.com/jinzhu/gorm"

type Alert struct {
	gorm.Model
	AlertId string      `json:"alert_id"`
	AlertReason  string `json:"alert_reason"`
	AlertMethod  string `json:"alert_method"`
	AlertOwner   string `json:"alert_owner"`
	AlertExOwner string `json:"alert_ex_owner"`
}