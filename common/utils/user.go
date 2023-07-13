package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// GenerateJwtToken 生成JWT
func GenerateJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// GeneratePassword 生成密码
func GeneratePassword(password string) string {
	pwd := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	return string(hashedPassword)
}

// ComparePassword 比较密码
func ComparePassword(password, hashedPassword string) bool {
	pwd := []byte(password)             //原密码
	hashedPwd := []byte(hashedPassword) //加密后的密码

	err := bcrypt.CompareHashAndPassword(hashedPwd, pwd) //注意参数顺序
	if err != nil {
		return false
	}
	return true
}
