package handlers

import (
	"learn-crud/models"
	"learn-crud/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	service services.SubjectService
}

func NewSubjectHandler(service services.SubjectService) *SubjectHandler {
	return &SubjectHandler{service: service}
}

// @Summary Update grade
// @Description Update an existing grade entry
// @Tags Nilai
// @Accept json
// @Produce json
// @Param id path int true "Grade ID"
// @Param nilai body models.Nilai true "Updated grade details"
// @Success 200 {object} models.Nilai
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /nilai/{id} [put]
func (h *SubjectHandler) GetAllSubjects(c *gin.Context) {
	subjects, err := h.service.GetAllSubjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subjects)
}

// @Summary Update grade
// @Description Update an existing grade entry
// @Tags Nilai
// @Accept json
// @Produce json
// @Param id path int true "Grade ID"
// @Param nilai body models.Nilai true "Updated grade details"
// @Success 200 {object} models.Nilai
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /nilai/{id} [put]
func (h *SubjectHandler) GetSubjectByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	subject, err := h.service.GetSubjectByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subject not found"})
		return
	}

	c.JSON(http.StatusOK, subject)
}

// @Summary Update grade
// @Description Update an existing grade entry
// @Tags Nilai
// @Accept json
// @Produce json
// @Param id path int true "Grade ID"
// @Param nilai body models.Nilai true "Updated grade details"
// @Success 200 {object} models.Nilai
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /nilai/{id} [put]
func (h *SubjectHandler) CreateSubject(c *gin.Context) {
	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newSubject, err := h.service.CreateSubject(subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newSubject)
}

func (h *SubjectHandler) UpdateSubject(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject.ID = uint(id)
	updatedSubject, err := h.service.UpdateSubject(subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedSubject)
}

func (h *SubjectHandler) DeleteSubject(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteSubject(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subject deleted successfully"})
}
