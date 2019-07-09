package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gosample/models"
	"gosample/util"
)

func ListStudent(c *gin.Context) {
	var students []models.Student
	err := models.GetAllStudent(&students)
	if err != nil {
		fmt.Println(err)
		util.RespondJSON(c, 404, students)
	} else {
		util.RespondJSON(c, 200, students)
	}
}


func GetById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	err := models.GetById(&student, id)
	if err != nil {
		util.RespondJSON(c, 404, student)
	} else {
		util.RespondJSON(c, 200, student)
	}
}


func AddStudent(c *gin.Context) {
	var student models.Student
	c.BindJSON(&student)
	err := models.Save(&student)
	if err != nil {
		util.RespondJSON(c, 404, student)
	} else {
		util.RespondJSON(c, 200, student)
	}
}

func DeleteById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	err := models.Delete(&student, id)
	if err != nil {
		util.RespondJSON(c, 404, student)
	} else {
		util.RespondJSON(c, 200, student)
	}
}