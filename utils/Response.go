package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func RespondJSON(w *gin.Context, status int, payload interface{}) {
	fmt.Println("status ", status)
	var res ResponseData

	res.Status = status
	//res.Meta = utils.ResponseMessage(status)
	res.Data = payload

	w.JSON(200, res)
}
