package utils

import (
  jwt "github.com/dgrijalva/jwt-go"
  "go-element-admin-api/configs"
  "strings"
  "time"
)

var jwtSecret = []byte(configs.ApplicationConfig.JwtSecret)

type Claims struct {
	UserId   int64   `json:"user_id"`
	UserName string  `json:"user_name"`
	Roles    []int64 `json:"roles"`
	jwt.StandardClaims
}

type JwtDo struct {
}


func (JwtDo) GenerateToken(userId int64, userName string, roles []int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
      userId,
      userName,
    roles,
      jwt.StandardClaims {
                  ExpiresAt : expireTime.Unix(),
                  Issuer : "bigfool",
		  },
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	var build strings.Builder
	build.WriteString("Bearer ")
	build.WriteString(token)
	token = build.String()
	return token, err
}

func (JwtDo) ParseToken(token string) (*Claims, error) {
	if strings.Contains(token,"Bearer ") {
		token = strings.Replace(token,"Bearer ","",1)
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

var Jwt = new(JwtDo)
