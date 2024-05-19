package common

import (
	"GinVue/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 加密秘钥
var jwtKey = []byte("a_secret_key")

// Claims 结构体
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// ReleaseToken 分发token
func ReleaseToken(user model.User) (string, error) {

	//过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			//token的过期时间
			ExpiresAt: expirationTime.Unix(),
			//token的签发时间
			IssuedAt: time.Now().Unix(),
			//签发者
			Issuer:  "jaehua",
			Subject: "user token",
		},
	}
	//new一个jwt类型的Token,指定签名算法和claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//将Token序列化为字符串并使用秘钥签名
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, err
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	//解析token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	// 检查是否过期
	if claims.ExpiresAt < time.Now().Unix() {
		return token, claims, errors.New("token expired")
	}

	return token, claims, err
}
