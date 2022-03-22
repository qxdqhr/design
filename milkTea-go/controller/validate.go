package controller

import (
	"fmt"
	"milkTea/service"
)

func registerValidate(user service.User)(error){
	//注册:用户输入手机号,密码,角色
	if len(user.Name)<=0{
		return fmt.Errorf("Name invalid")
	}
	if len(user.Password)<6 || len(user.Password)>16 {
		//密码长度小于0且大于16
		return fmt.Errorf("Password invalid")
	}
	if len(user.Telephone)!=11 {
		//手机号长度小于0且大于16
		return fmt.Errorf("Telephone invalid")
	}
	if !(user.Role=="Owner" || user.Role=="ExOwner"){
		return fmt.Errorf("Role invalid")
	}
	return nil;
}
func loginValidate(user service.User)(error){

	if len(user.Userid)!=11 {
		return fmt.Errorf("Userid invalid")
	}
	if len(user.Password)<6 || len(user.Password)>16 {
		//密码长度小于0且大于16
		return fmt.Errorf("Password invalid")
	}
	if !(user.Role=="Owner" || user.Role=="ExOwner"){
		//角色长度小于0且大于16
		return fmt.Errorf("Role invalid")
	}
	return nil;
}

func orderValidate(order service.Order)(error){
	if len(order.CustomerId)<=0 {
		fmt.Errorf("username invalid")
	}
	if len(order.CustomerId)<=0 {
		fmt.Errorf("username invalid")
	}
	return nil
}