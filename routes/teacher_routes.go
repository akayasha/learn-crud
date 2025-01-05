package routes

import (
	"learn-crud/handlers"

	"github.com/gin-gonic/gin"
)

func TeacherRoutes(router *gin.Engine, handler *handlers.TeacherHandler) {
	router.GET("/teachers", handler.GetAllTeachers)
	router.GET("/teachers/:nip", handler.GetTeacherByNIP)
	router.POST("/teachers", handler.CreateTeacher)
	router.PUT("/teachers/:nip", handler.UpdateTeacher)
	router.DELETE("/teachers/:nip", handler.DeleteTeacher)
	router.GET("/teachers/status", handler.GetTeachersByStatus)
}
