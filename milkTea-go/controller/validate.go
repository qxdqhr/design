package controller

import (
	"fmt"
	"milkTea/service"
	"strconv"
	"strings"
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
		return fmt.Errorf("consumerName invalid")
	}
	if len(order.Buyingjuice)<=0 {
		return fmt.Errorf("Buyingjuice invalid")
	}
	if len(order.JuiceNumber)<=0 {
		return fmt.Errorf("JuiceNumber invalid")
	}
	if len(order.OrderingTime)<=0 {
		return fmt.Errorf("OrderingTime invalid")
	}
	if len(order.TotalSellingPrice)<=0 {
		return fmt.Errorf("TotalSellingPrice invalid")
	}
	if !(order.CurEvaluate == "好评" ||order.CurEvaluate == "中评" || order.CurEvaluate == "差评") {
		return fmt.Errorf("Buyingjuice invalid")
	}
	return nil
}
func orderQueryValidate(query *service.Query)(error){
	fmt.Print(query)
	if query.Func != "order" {
		return fmt.Errorf("func failed")
	}
	switch query.QueryName{
	case  "顾客名":
		if len(query.QueryValue) == 0 {
			return fmt.Errorf("顾客名称错误")
		}
		break
	case  "购买饮品":
		times := strings.Split(query.QueryValue,"｜")
		if len(times) == 0 {
			return fmt.Errorf("饮品格式错误")
		}
		for _,val := range times{
			juiceToNum := strings.Split(val,":")
			if len(juiceToNum) == 0 {
				return fmt.Errorf("饮品格式错误")
			}
			if juiceToNum[0]=="" {
				return fmt.Errorf("饮品格式错误")
			}
			if _,err := strconv.ParseInt(juiceToNum[1],10,64);err!=nil{
				return fmt.Errorf("饮品格式错误"+err.Error())
			}
		}
		break
	case  "下单时间":
		times := strings.Split(query.QueryValue,"/")
		if len(times) == 0 {
			return fmt.Errorf("时间格式错误")
		}
		for _,val := range times{
			_,err := strconv.ParseInt(val,10,64)
			if err!=nil {
				return fmt.Errorf("时间格式错误")
			}
		}
		break
	case  "本单饮品数量":
		if  _,err := strconv.ParseInt(query.QueryValue,10,64);err!=nil {
			return fmt.Errorf("数量格式错误")
		}
		break
	case  "总售价":
		if  _,err := strconv.ParseInt(query.QueryValue,10,64);err!=nil {
			return fmt.Errorf("价格格式错误")
		}
		break
	case "本单评价":
		if !(query.QueryValue == "好评" ||query.QueryValue == "中评" || query.QueryValue == "差评"){
			return fmt.Errorf("查询格式错误")
		}
		break
	case "" :
		return fmt.Errorf("查询格式错误")
		break
	}
	return nil
}