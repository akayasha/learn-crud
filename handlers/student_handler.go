package handlers

import (
	"learn-crud/models"
	"learn-crud/services"
	"learn-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	service services.StudentService
}

func NewStudentHandler(service services.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

// @Summary Get all students
// @Description Retrieve a list of all students from the database
// @Tags Student
// @Accept json
// @Produce json
// @Success 200 {array} models.Student
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /students [get]
func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.service.GetAllStudents()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to retrieve students")
		return
	}
	utils.Respond(c, http.StatusOK, "Success", students)
}

// @Sumary Get Students by kelas
// @Description Retrieve a list of al student by kelas
// @Tags Student
// @Accept json
// @Produce json
// @Param nip path string true "student kelas"
// @Success 200 {array} models.Student
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /students/kelas [get]
func (h *StudentHandler) GetStudentByKelas(c *gin.Context) {
	kelas := c.DefaultQuery("kelas", "")
	if kelas == "" {
		utils.RespondError(c, http.StatusBadRequest, "Kelas Query is required")
	}
	students, err := h.service.GetStudentByKelas(kelas)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to retrieve students")
		return
	}
	utils.Respond(c, http.StatusOK, "Success", students)
}

// @Sumary Get Students by name
// @Description Retrieve a list of al student by name
// @Tags Student
// @Accept json
// @Produce json
// @Param nip path string true "Student name"
// @Success 200 {array} models.Student
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /students/name [get]
func (h *StudentHandler) GetStudentByName(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	if name == "" {
		utils.RespondError(c, http.StatusBadRequest, "Name query parameter is required")
		return
	}
	student, err := h.service.GetStudentByName(name)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to retrieve student")
		return
	}
	utils.Respond(c, http.StatusOK, "Success", student)
}

// @Sumary Get Students by nim
// @Description Retrieve a list of al student by nim
// @Tags Student
// @Accept json
// @Produce json
// @Param nip path string true "Student nim"
// @Success 200 {array} models.Student
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /students/nim [get]
func (h *StudentHandler) GetStudentByNIM(c *gin.Context) {
	nim := c.DefaultQuery("nim", "")
	if nim == "" {
		utils.RespondError(c, http.StatusBadRequest, "NIM query parameter is required")
		return
	}
	student, err := h.service.GetStudentByNIM(nim)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Student not found")
		return
	}
	utils.Respond(c, http.StatusOK, "Success", student)
}

// @Sumary Get Students by nim
// @Description Retrieve a list of al student by status
// @Tags Student
// @Accept json
// @Produce json
// @Param nip path string true "Student status"
// @Success 200 {array} models.Student
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /students/status [get]
func (h *StudentHandler) GetStudentByStatus(c *gin.Context) {
	status := c.DefaultQuery("status", "")
	if status == "" {
		utils.RespondError(c, http.StatusBadRequest, "Status query parameter is required")
		return
	}
	student, err := h.service.GetStudentByStatus(status)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to retrieve student")
		return
	}
	utils.Respond(c, http.StatusOK, "Success", student)

}

// @Summary Create a new student
// @Description Add a new student to the database
// @Tags student
// @Accept json
// @Produce json
// @Param student body models.Student true "student object"
// @Success 201 {object} models.Student
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /student [post]
func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student

	// Set default status if it's empty
	if student.Status == "" {
		student.Status = models.Active
	}

	// Validate JSON body and bind it to the Student struct
	if err := c.ShouldBindJSON(&student); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	// Additional runtime validation for empty or null fields
	missingFields := utils.ValidateStruct(student)
	if missingFields != "" {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request: "+missingFields)
		return
	}

	createdStudent, err := h.service.CreateStudent(student)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create student")
		return
	}
	utils.Respond(c, http.StatusCreated, "Student created successfully", createdStudent)
}

// @Summary Update student
// @Description Update an existing student's details
// @Tags Student
// @Accept json
// @Produce json
// @Param nim path string true "Student NIM"
// @Param student body models.Student true "Student object"
// @Success 200 {object} models.Student
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /student/{nim} [put]
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	nim := c.Param("nim") // Use nim instead of id
	var student models.Student

	// Validate JSON body and bind it to the Student struct
	if err := c.ShouldBindJSON(&student); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	// Additional runtime validation for empty or null fields
	missingFields := utils.ValidateStruct(student)
	if missingFields != "" {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request: "+missingFields)
		return
	}

	student.NIM = nim
	updatedStudent, err := h.service.UpdateStudent(student)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update student")
		return
	}

	utils.Respond(c, http.StatusOK, "Student updated successfully", updatedStudent)
}

// @Summary Delete student
// @Description Remove a student by their NIM
// @Tags Student
// @Accept json
// @Produce json
// @Param nim path string true "Student NIM"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /student/{nim} [delete]
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	nim := c.Param("nim") // Use nim instead of id
	err := h.service.DeleteStudent(nim)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete student")
		return
	}
	utils.Respond(c, http.StatusOK, "Student deleted successfully", nil)
}
