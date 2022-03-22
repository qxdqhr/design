package dto

import (
	"milkTea/service"
)

type DTOLoginAuthorize struct {
	Userid    string `json:"userid"`
	Role      string `json:"role"`
}
func ToReturnLoginAuthorize(u interface{}) DTOLoginAuthorize {
	user,_:=u.(*service.User)
	return DTOLoginAuthorize{user.Userid,user.Role}
}
