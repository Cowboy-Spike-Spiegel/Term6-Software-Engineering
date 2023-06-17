package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// jwt加密密钥
var JwtKey = []byte("Spike")

// token的claim
type Claims struct {
	UserId string
	Auth   int
	jwt.StandardClaims
}

// 发放token
func ReleaseToken(user_id string, auth int) (string, error) {
	//token的有效期
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		//自定义字段
		UserId: user_id,
		Auth:   auth,
		//标准字段
		StandardClaims: jwt.StandardClaims{
			//过期时间
			ExpiresAt: expirationTime.Unix(),
			//发放的时间
			IssuedAt: time.Now().Unix(),
			//发放者
			Issuer: "Spike",
			//主题
			Subject: "user token",
		},
	}

	//使用jwt密钥生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	//返回token
	return tokenString, nil
}

// 从tokenString中解析出claims并返回
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})
	return token, claims, err
}
