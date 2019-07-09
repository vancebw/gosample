package controllers

import (
	"gosample/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gosample/constants"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := "test"
	password := "test"
	token, err := util.GenerateToken(username, password)
	if err != nil {
		util.RespondJSON(c, http.StatusInternalServerError, nil)
		return
	}

	util.RespondJSON(c, constants.SUCCESS, map[string]string{
		"token": token,
	})
}
