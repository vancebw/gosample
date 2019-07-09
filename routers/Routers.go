package routers

import (
	"github.com/gin-gonic/gin"
	"gosample/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("students", controllers.ListStudent)
		v1.GET("student", controllers.GetById)
		v1.PUT("student", controllers.AddStudent)
		v1.DELETE("student", controllers.DeleteById)
	}

	return r
}
