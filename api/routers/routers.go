package routers

import (
	"github.com/gin-gonic/gin"
	"gosample/api/controllers"
	"gosample/api/util"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/auth", controllers.GetAuth)
	v1 := r.Group("/v1")
	v1.Use(util.JWT())
	{
		v1.GET("students", controllers.ListStudent)
		v1.GET("student", controllers.GetById)
		v1.PUT("student", controllers.AddStudent)
		v1.DELETE("student", controllers.DeleteById)
	}

	return r
}
