package routes

import (
	"learn-crud/handlers"

	"github.com/gin-gonic/gin"
)

func SubjectRoutes(router *gin.Engine, handler *handlers.SubjectHandler) {
	router.GET("/subjects", handler.GetAllSubjects)
	router.GET("/subjects/:id", handler.GetSubjectByID)
	router.POST("/subjects", handler.CreateSubject)
	router.PUT("/subjects/:id", handler.UpdateSubject)
	router.DELETE("/subjects/:id", handler.DeleteSubject)
}
