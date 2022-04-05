package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/spf13/viper"
)


func GetDB()(* gorm.DB){
	//driverName	:=viper.GetString("datasource.driverName")
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

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})


	//,args)

	if err!= nil{
		Err(err.Error())
	}

	return db
}
