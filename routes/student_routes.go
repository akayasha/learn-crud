package routes

import (
	"learn-crud/handlers"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.Engine, handler *handlers.StudentHandler) {
	router.GET("/students", handler.GetStudents)
	router.GET("/students/nim", handler.GetStudentByNIM)
	router.POST("/students", handler.CreateStudent)
	router.PUT("/students/:nim", handler.UpdateStudent)
	router.DELETE("/students/:nim", handler.DeleteStudent)
	router.GET("/students/name", handler.GetStudentByName)
	router.GET("/students/kelas", handler.GetStudentByKelas)
	router.GET("/students/status", handler.GetStudentByStatus)
}
