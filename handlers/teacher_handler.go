package handlers

import (
	"learn-crud/models"
	"learn-crud/services"
	"learn-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeacherHandler struct {
	service services.TeacherService
}

func NewTeacherHandler(service services.TeacherService) *TeacherHandler {
	return &TeacherHandler{service: service}
}

// @Summary Get all teachers
// @Description Retrieve all teachers from the database
// @Tags Teacher
// @Accept json
// @Produce json
// @Success 200 {array} models.Teacher
// @Failure 500 {object} gin.H
// @Router /teachers [get]
func (h *TeacherHandler) GetAllTeachers(c *gin.Context) {
	teachers, err := h.service.GetAllTeachers()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Respond(c, http.StatusOK, "Success", teachers)
}

// @Summary Get teacher by NIP
// @Description Retrieve a teacher by their NIP
// @Tags Teacher
// @Accept json
// @Produce json
// @Param nip path string true "Teacher NIP"
// @Success 200 {object} models.Teacher
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /teachers/{nip} [get]
func (h *TeacherHandler) GetTeacherByNIP(c *gin.Context) {
	nip := c.Param("nip")
	if nip == "" {
		utils.RespondError(c, http.StatusBadRequest, "NIP is required")
		return
	}

	teacher, err := h.service.GetTeacherByNIP(nip)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Respond(c, http.StatusOK, "Success", teacher)
}

// @Summary Create a new teacher
// @Description Add a new teacher to the database
// @Tags Teacher
// @Accept json
// @Produce json
// @Param teacher body models.Teacher true "Teacher object"
// @Success 201 {object} models.Teacher
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /teachers [post]
func (h *TeacherHandler) CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	newTeacher, err := h.service.CreateTeacher(teacher)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Respond(c, http.StatusCreated, "Teacher created successfully", newTeacher)
}

// @Summary Update teacher
// @Description Update an existing teacher's details
// @Tags Teacher
// @Accept json
// @Produce json
// @Param nip path string true "Teacher NIP"
// @Param teacher body models.Teacher true "Teacher object"
// @Success 200 {object} models.Teacher
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /teachers/{nip} [put]
func (h *TeacherHandler) UpdateTeacher(c *gin.Context) {
	nip := c.Param("nip")
	if nip == "" {
		utils.RespondError(c, http.StatusBadRequest, "NIP is required")
		return
	}

	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	teacher.NIP = nip
	updatedTeacher, err := h.service.UpdateTeacher(teacher)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Respond(c, http.StatusOK, "Teacher updated successfully", updatedTeacher)
}

// @Summary Delete teacher
// @Description Remove a teacher by their NIP
// @Tags Teacher
// @Accept json
// @Produce json
// @Param nip path string true "Teacher NIP"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /teachers/{nip} [delete]
func (h *TeacherHandler) DeleteTeacher(c *gin.Context) {
	nip := c.Param("nip")
	if nip == "" {
		utils.RespondError(c, http.StatusBadRequest, "NIP is required")
		return
	}

	err := h.service.DeleteTeacher(nip)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Respond(c, http.StatusOK, "Teacher deleted successfully", nil)
}

// @Summary Get teachers by status
// @Description Retrieve teachers by their status
// @Tags Teacher
// @Accept json
// @Produce json
// @Param status query string true "Teacher status"
// @Success 200 {array} models.Teacher
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /teachers/status [get]
func (h *TeacherHandler) GetTeachersByStatus(c *gin.Context) {
	status := c.Query("status")
	if status == "" {
		utils.RespondError(c, http.StatusBadRequest, "Status is required")
		return
	}

	teachers, err := h.service.GetTeachersByStatus(status)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Respond(c, http.StatusOK, "Success", teachers)
}
