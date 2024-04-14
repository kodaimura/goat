package jwt

import (
	"time"
	"encoding/json"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	jwtpackage "github.com/golang-jwt/jwt/v4"

	"goat/config"
)


func SetTokenToCookie (c *gin.Context, pl JwtPayload) error {
	jwtStr, err := EncodeJwt(pl)
	if err != nil {
		return err
	}
	cf := config.GetConfig()
	c.SetCookie(COOKIE_KEY_JWT, jwtStr, int(JWT_EXPIRES), "/", cf.AppHost, false, true)
	return nil
}


func RemoveTokenFromCookie (c *gin.Context) {
	cf := config.GetConfig()
	c.SetCookie(COOKIE_KEY_JWT, "", 0, "/", cf.AppHost, false, true)
}


func GetPayload(c *gin.Context) JwtPayload {
	pl := c.Keys[CONTEXT_KEY_PAYLOAD]
	if pl == nil {
		return JwtPayload{}
	}
	return pl.(JwtPayload)
}


func EncodeJwt (pl JwtPayload) (string, error) {
	return encodeJwt(pl)
}

func ExpireJwt (pl JwtPayload) JwtPayload {
	pl.IssuedAt =  time.Now().Unix()
	pl.ExpiresAt = time.Now().Unix()
	return pl
}  


func Auth (c *gin.Context) error {
	tokenStr, err := getJwtToken(c)
	if err != nil {
		return err
	}

	pl, err := decodeJwt(tokenStr)
	if err != nil {
		return err
	}
	
	c.Set(CONTEXT_KEY_PAYLOAD, pl)
	return nil
}


func encodeJwt (pl JwtPayload) (string, error) {
	cf := config.GetConfig()
	token := jwtpackage.NewWithClaims(jwtpackage.SigningMethodHS256, pl)
	return token.SignedString([]byte(cf.JwtSecretKey))
}


func decodeJwt (encoded string) (JwtPayload, error) {
	cf := config.GetConfig()
	token, err := jwtpackage.Parse(encoded, func(token *jwtpackage.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtpackage.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(cf.JwtSecretKey), nil
	})
	if err != nil {
		return JwtPayload{}, err
	}

	return convertToPayload(token)
}


func getJwtToken (c *gin.Context) (string, error) {
	token, err := c.Cookie(COOKIE_KEY_JWT)
	if err == nil {
		return token, nil
	}

	bearer := c.Request.Header.Get("Authorization")
	if bearer != "" {
		if strings.Index(bearer, "Bearer ") != 0 {
			return strings.TrimSpace(bearer[7:]), nil
		}
	}

	return "", errors.New("Token not found")
}


func convertToPayload (token *jwtpackage.Token) (JwtPayload, error) {
	var pl JwtPayload

	jsonString, err := json.Marshal(token.Claims.(jwtpackage.MapClaims))

	if err == nil {
		err = json.Unmarshal(jsonString, &pl)
	}

	return pl, err
}