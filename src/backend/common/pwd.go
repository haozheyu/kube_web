package common

import (
	"github.com/beego/beego/v2/adapter/logs"
	"golang.org/x/crypto/bcrypt"
)

//加密
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd,bcrypt.MinCost)
	if err != nil {
		logs.Error(err.Error())
	}
	return string(hash)
}

//密码验证
func ParsPwd(hashPwd string,plainPwd []byte) bool {
	byteHash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash,plainPwd) //加密的密码和原来的密码对比
	if err != nil {
		return false
	}
	return true
}
