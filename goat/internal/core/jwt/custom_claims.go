package jwt

import (
	"errors"

	"github.com/gin-gonic/gin"
)

/*
JwtPayload拡張
*/

type CustomClaims struct {
	UserId int
	UserName string
}


func GetUserId (c *gin.Context) (int, error) {
	pl := c.Keys[CONTEXT_KEY_PAYLOAD]
	if pl == nil {
		return -1, errors.New("GetUserId error")
	} else {
		return pl.(JwtPayload).UserId, nil
	}	
}

func GetUserName (c *gin.Context) (string, error) {
	pl := c.Keys[CONTEXT_KEY_PAYLOAD]
	if pl == nil {
		return "", errors.New("GetUserName error")
	} else {
		return pl.(JwtPayload).UserName, nil
	}
}