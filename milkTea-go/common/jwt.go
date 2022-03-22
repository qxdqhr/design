package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)
var jwtKey = []byte("secretword")

type Claims struct {
	UserId string
	jwt.StandardClaims
}
//发放token
func CreateToken(userid string) (string,error){//创建token并返回的函数
	//toke 的有效时间
	expireTime := time.Now().Add(7*24*time.Hour)
	claims := Claims{
		UserId:         userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "qhr",//发放者
			Subject:   "user token",//主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err :=token.SignedString(jwtKey)
	if err!=nil{
		return "",err
	}
	return tokenString,nil;
}
//解析token
func AnalyzeToken(tokenString string)(*jwt.Token,*Claims,error){

	claims:=&Claims{}
	token, err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey,nil;
	})
	if err != nil {
		return nil,nil,err
	}
	return token,claims,nil
}
