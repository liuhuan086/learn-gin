package api

import (
	"example/models"
	"example/pkg/e"
	"example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required"`
	Password string `valid:"Required"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}

	au := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&au)

	data := make(map[string]interface{})
	code := e.InvalidParams
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ErrorAuthToken
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ErrorAuth
		}
	} else {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "参数：" + err.Key + "错误， " + err.Message,
				"data": make(map[string]string),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
