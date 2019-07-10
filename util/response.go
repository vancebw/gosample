package util

import (
	"github.com/gin-gonic/gin"
	"gosample/constants"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RespondJSON(status int, w *gin.Context, code int, payload interface{}) {
	var res ResponseData
	res.Code = code
	res.Msg = constants.GetMsg(code)
	res.Data = payload

	w.JSON(status, res)
}
