package routes

import (
	"learn-crud/handlers"

	"github.com/gin-gonic/gin"
)

func NilaiRoutes(router *gin.Engine, handler *handlers.NilaiHandler) {
	router.GET("/nilai", handler.GetNilaiByNIM)
	router.POST("/nilai", handler.CreateNilai)
	router.PUT("/nilai/:id", handler.UpdateNilai)
	router.DELETE("/nilai/:id", handler.DeleteNilai)
}
