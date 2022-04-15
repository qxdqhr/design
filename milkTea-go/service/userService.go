package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"milkTea/common"
)
type User struct {
	Name      string `json:"name"`
	Userid    string `json:"user_id" gorm:"column:user_id"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Telephone string `json:"telephone"`
	ExOwnerID string `json:"exownerid"" gorm:"column:ex_owner_id"`
}
func LoginService(user *User)(error,*User){
	db:=common.GetDB()
	db.AutoMigrate(&User{})

	u:=new(User)
	dbq:=db.First(&u, "user_id = ?", user.Userid)
	fmt.Println(u)

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
	db.AutoMigrate(&User{})
	//判断当前uid是否已存在
	u:=new(User)
	dbq:=db.Find(&u, "telephone = ?", user.Telephone)
	if dbq.RowsAffected>0{
		return fmt.Errorf("此电话号已经注册过")
	}else if u.Userid == user.Userid {
		return fmt.Errorf("此电话号已经注册过")
	}else{
		db.Create(&user)
		return nil
	}
}
func GetAllSubOwner(userid string)([]string,error){
	db:=common.GetDB()
	db.AutoMigrate(&User{})
	//判断当前uid是否已存在
	u:=make([]User,0)
	dbq:=db.Debug().Find(&u, "ex_owner_id = ?",userid)
	if dbq.RowsAffected<0{
		fmt.Println("此经销商没有加盟商")
		return nil,nil
	}
	fmt.Println(u)
	res:=make([]string,0)

	for _,v:=range u{
		res = append(res, v.Userid)
	}
	return res,nil
}
