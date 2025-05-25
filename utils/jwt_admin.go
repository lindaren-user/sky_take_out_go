package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var jwtKey = []byte(viper.GetString("jwt.admin.key"))

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// 生成 Token
func GenerateJWT(userId int) (string, error) {
	claims := Claims{
		UserID: userId, // 注意 claims，这里面自己有ID
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)), // 2小时过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "lwj",
		},
	}

	// 创建一个新的 token 实例，指定使用 HMAC-SHA256 签名算法，并将 claims 作为它的负载内容。
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用预设的密钥 jwtKey 对 token 进行签名，并返回最终的字符串。
	return token.SignedString(jwtKey)
}

// 解析 Token
func ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// 解析出的 Claims 类型断言成功，token 是有效的（签名正确、未过期等）。
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
