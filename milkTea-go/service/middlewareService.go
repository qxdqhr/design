package service

import "milkTea/common"

func LoginAuthorizeService(user *User)(*User){
	db:=common.GetDB()
	db.AutoMigrate(&User{})
	u:=User{}
	db.First(&u,"user_id = ?",user.Userid)
	if u.Userid=="" {
		//查无此人
		return nil
	}
	return &u
}

