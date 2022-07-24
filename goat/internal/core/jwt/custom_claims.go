package jwt

import (
	"errors"

	"github.com/gin-gonic/gin"
)

/*
JwtPayload拡張
*/

type CustomClaims struct {
	UserId int `json:"userid"`
	Username string `json:"username"`
}


func GetUserId (c *gin.Context) (int, error) {
	pl := c.Keys[CONTEXT_KEY_PAYLOAD]
	if pl == nil {
		return -1, errors.New("GetUserId error")
	} else {
		return pl.(JwtPayload).UserId, nil
	}	
}

func GetUsername (c *gin.Context) (string, error) {
	pl := c.Keys[CONTEXT_KEY_PAYLOAD]
	if pl == nil {
		return "", errors.New("GetUsername error")
	} else {
		return pl.(JwtPayload).Username, nil
	}
}