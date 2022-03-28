package service

import (
	"github.com/jinzhu/gorm"
)

type Owner struct {
	gorm.Model `json:"-"`
	Userid    string `json:"userid"`
	AlertTimes int `json:"alert_times"`
	RecentReason string `json:"recent_reason"`
	ExOwnerUserid string `json:"exowneruserid"`
}







