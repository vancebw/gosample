package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gosample/constants"
	"gosample/models"
	"gosample/util"
	"net/http"
)

func ListStudent(c *gin.Context) {
	var students []models.Student
	err := models.GetAllStudent(&students)
	if err != nil {
		fmt.Println(err)
		util.RespondJSON(http.StatusOK, c, constants.NOT_FOUND, students)
	} else {
		util.RespondJSON(http.StatusOK, c, constants.SUCCESS, students)
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
