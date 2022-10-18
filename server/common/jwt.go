package common

import (
	"github.com/golang-jwt/jwt"
	"imall/global"
	"time"
)

var SigningKey = []byte(global.Config.Jwt.SigningKey)

type Claims struct {
	Username string `json:"username"`
	OpenId   string `json:"openId"`
	jwt.StandardClaims
}

// GenerateToke 生成Token
func GenerateToke(username string) (string, error) {
	claims := Claims{Username: username, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 60*60,
		Issuer:    username,
	},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SigningKey)
}

// GenerateTokeOpenId 生成Token
func GenerateTokeOpenId(openId string) (string, error) {
	claims := Claims{OpenId: openId, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 60*60,
		Issuer:    openId,
	},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SigningKey)
}

// VerifyToken 验证Token
func VerifyToken(tokenString string) error {
	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	return err
}

// ParseToken 验证Token
func ParseToken(tokenString string) (*Claims, error) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, nil
}
