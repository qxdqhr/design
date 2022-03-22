package config

import (
	"github.com/spf13/viper"
	"milkTea/common"
	"os"
)

//使用viper 作为config组件
func InitConfig(){
	workdir,_ := os.Getwd()//返回当前目录
	viper.SetConfigName("application")//配置文件名
	viper.SetConfigType("yml")//配置文件类型
	viper.AddConfigPath(workdir+"/common/config")//配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		common.Err("configuration read failed")
		return
	}
}
