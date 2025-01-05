package handlers

import (
	"learn-crud/models"
	"learn-crud/services"
	"learn-crud/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NilaiHandler struct {
	service services.NilaiService
}

func NewNilaiHandler(service services.NilaiService) *NilaiHandler {
	return &NilaiHandler{service: service}
}

// @Summary Get grades by NIM
// @Description Retrieve all grades for a student by their NIM
// @Tags Nilai
// @Accept json
// @Produce json
// @Param nim query string true "Student NIM"
// @Success 200 {array} models.Nilai
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /nilai [get]
func (h *NilaiHandler) GetNilaiByNIM(c *gin.Context) {
	nim := c.Query("nim")
	if nim == "" {
		utils.RespondError(c, http.StatusBadRequest, "NIM query parameter is required")
		return
	}

	nilai, err := h.service.GetNilaiByNIM(nim)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to retrieve grades: "+err.Error())
		return
	}

	utils.Respond(c, http.StatusOK, "Success", nilai)
}

// @Summary Get grades by NIM
// @Description Retrieve all grades for a student by their NIM
// @Tags Nilai
// @Accept json
// @Produce json
// @Param nim query string true "Student NIM"
// @Success 200 {array} models.Nilai
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /nilai [get]
func (h *NilaiHandler) CreateNilai(c *gin.Context) {
	var input models.Nilai
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	newNilai, err := h.service.CreateNilai(input)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Respond(c, http.StatusCreated, "Grade created successfully", newNilai)
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
func (h *NilaiHandler) UpdateNilai(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	var nilai models.Nilai
	if err := c.ShouldBindJSON(&nilai); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	nilai.ID = uint(id)
	updatedNilai, err := h.service.UpdateNilai(nilai)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update grade: "+err.Error())
		return
	}

	utils.Respond(c, http.StatusOK, "Grade updated successfully", updatedNilai)
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
func (h *NilaiHandler) DeleteNilai(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	err = h.service.DeleteNilai(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete grade: "+err.Error())
		return
	}

	utils.Respond(c, http.StatusOK, "Grade deleted successfully", nil)
}
