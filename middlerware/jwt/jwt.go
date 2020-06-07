package jwt

import (
	"example/pkg/e"
	"example/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == ""{
			code = e.ErrorUnauthorized
		} else {
			claims, err := util.ParseToken(token)
			if err != nil{
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}

		if code != e.SUCCESS{
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}