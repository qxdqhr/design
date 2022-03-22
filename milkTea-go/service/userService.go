package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"milkTea/common"
)
type User struct {
	gorm.Model
	Name    string `json:"name"`
	Userid    string `json:"userid"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Telephone string `json:"telephone"`
}
func LoginService(user *User)(error,*User){
	db:=common.GetDB()
	db.AutoMigrate(&User{})

	u:=new(User)
	dbq:=db.First(&u, "userid = ?", user.Userid)
	fmt.Println(u)
	defer db.Close()

	if dbq.RowsAffected<1{//没查到
		err := fmt.Errorf("用户工号不存在")
		return err,nil
	}
	if err:=bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(user.Password));err != nil{
		//密码错误
		err := fmt.Errorf("密码错误")

		return err,nil
	}
	return nil,u

}

func RegisterService(user *User)(error){
	db:=common.GetDB()
	db=db.AutoMigrate(&User{})
	//判断当前uid是否已存在
	u:=new(User)
	dbq:=db.Find(&u, "telephone = ?", user.Telephone)
	defer db.Close()
	if dbq.RowsAffected>0{
		return fmt.Errorf("此电话号已经注册过")
	}else if u.Userid == user.Userid {
		return fmt.Errorf("此电话号已经注册过")
	}else{
		db.Create(&user)
		return nil
	}
}
