package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"gosample/api/constants"
	"gosample/api/util"
	"gosample/portal/models"
	proto "gosample/portal/proto"
	"net/http"
)

func ListStudent(c *gin.Context) {
	resp, err := proto.NewStudentService("gosample.srv.portal", client.DefaultClient).ListStudent(context.TODO(), &proto.Request{})
	if err != nil {
		util.RespondJSON(http.StatusOK, c, constants.NOT_FOUND, err)
	} else {
		util.RespondJSON(http.StatusOK, c, constants.SUCCESS, resp)
	}
}

func GetById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	err := models.GetById(&student, id)
	if err != nil {
		util.RespondJSON(http.StatusOK, c, constants.NOT_FOUND, student)
	} else {
		util.RespondJSON(http.StatusOK, c, constants.SUCCESS, student)
	}
}

func AddStudent(c *gin.Context) {
	var student models.Student
	c.BindJSON(&student)
	err := models.Save(&student)
	if err != nil {
		util.RespondJSON(http.StatusOK, c, constants.NOT_FOUND, student)
	} else {
		util.RespondJSON(http.StatusOK, c, constants.SUCCESS, student)
	}
}

func DeleteById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	err := models.Delete(&student, id)
	if err != nil {
		util.RespondJSON(http.StatusOK, c, constants.NOT_FOUND, student)
	} else {
		util.RespondJSON(http.StatusOK, c, constants.SUCCESS, student)
	}
}
