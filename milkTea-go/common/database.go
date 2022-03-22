package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)


func GetDB()(* gorm.DB){
	driverName	:=viper.GetString("datasource.driverName")
	userName	:=viper.GetString("datasource.userName")
	password	:=viper.GetString("datasource.password")
	host		:=viper.GetString("datasource.host")
	port		:=viper.GetString("datasource.port")
	DBName		:=viper.GetString("datasource.DBName")
	charset		:=viper.GetString("datasource.charset")
	args:=fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		userName,
		password,
		host,
		port,
		DBName,
		charset,
	)

	DB,err := gorm.Open(driverName,args)
	DB.SingularTable(true)
	if err!= nil{
		Err(err.Error())
	}

	return DB
}
