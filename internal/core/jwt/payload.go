package jwt

import (
	"time"
	
	jwtpackage "github.com/golang-jwt/jwt/v4"
)


type JwtPayload struct {
	jwtpackage.StandardClaims
	CustomClaims
}

type CustomClaims struct {
	UserId int
	UserName string
	/* 独自のフィールドを追加可能 */
}


func NewPayload(claims CustomClaims) JwtPayload {
	var pl JwtPayload

	pl.CustomClaims = claims
	pl.IssuedAt =  time.Now().Unix()
	pl.ExpiresAt = time.Now().Add(time.Second * JWT_EXPIRES).Unix()

	return pl
}