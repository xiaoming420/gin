package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"gin/pkg/setting"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	UserId   int
	jwt.StandardClaims
}

func GenerateToken(user_id  int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		user_id,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "gin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}