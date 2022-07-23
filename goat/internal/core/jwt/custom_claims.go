package jwt

import (
	"errors"

    "github.com/gin-gonic/gin"
)

/*
JwtPayload拡張
*/

type CustomClaims struct {
	UId int `json:"uid"`
    Username string `json:"username"`
}


func GetUId (c *gin.Context) (int, error) {
	pl := c.Keys[CONTEXT_KEY_PAYLOAD]
	if pl == nil {
		return -1, errors.New("GetUId error")
	} else {
		return pl.(JwtPayload).UId, nil
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