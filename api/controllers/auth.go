package controllers

import (
	"github.com/gin-gonic/gin"
	"gosample/api/constants"
	"gosample/api/util"
	"net/http"
)

func GetAuth(c *gin.Context) {
	username := "test"
	password := "test"
	token, err := util.GenerateToken(username, password)
	if err != nil {
		util.RespondJSON(http.StatusOK, c, constants.ERROR, nil)
		return
	}

	util.RespondJSON(http.StatusOK, c, constants.SUCCESS, map[string]string{
		"token": token,
	})
}
