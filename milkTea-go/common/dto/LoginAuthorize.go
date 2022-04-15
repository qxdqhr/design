package dto

import (
	"milkTea/service"
)

type DTOLoginAuthorize struct {
	User_id    string `json:"user_id"`
	Role      string `json:"role"`
}
func ToReturnLoginAuthorize(u interface{}) DTOLoginAuthorize {
	user,_:=u.(*service.User)
	return DTOLoginAuthorize{user.Userid,user.Role}
}
