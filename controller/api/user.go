package api

import (
	"log"
	"net/http"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"gin/models"
	"gin/pkg/e"
	"gin/pkg/util"
)

type auth struct {
	Phone string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	phone := c.PostForm("phone")

	password := c.PostForm("password")

	valid := validation.Validation{}
	a := auth{Phone: phone, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		user_id := models.CheckAuth(phone, password)
		if user_id > 0 {
			token, err := util.GenerateToken(user_id)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
