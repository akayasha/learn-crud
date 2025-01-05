package repository

import (
	"gorm.io/gorm"
	"learn-crud/models"
)

type TeacherRepository interface {
	FindAll() ([]models.Teacher, error)
	FindByNip(nip string) (models.Teacher, error)
	CreateTeacher(teacher models.Teacher) (models.Teacher, error)
	UpdateTeacher(teacher models.Teacher) (models.Teacher, error)
	DeleteTeacher(nip string) error
	FindByStatus(status string) ([]models.Teacher, error)
}

type teacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &teacherRepository{db: db}
}

func (r *teacherRepository) FindAll() ([]models.Teacher, error) {
	var teachers []models.Teacher
	err := r.db.Find(&teachers).Error
	return teachers, err
}

func (r *teacherRepository) FindByNip(nip string) (models.Teacher, error) {
	var teacher models.Teacher
	err := r.db.Where("nip = ? ", nip).First(&teacher).Error
	return teacher, err
}

func (r *teacherRepository) CreateTeacher(teacher models.Teacher) (models.Teacher, error) {
	err := r.db.Create(&teacher).Error
	return teacher, err
}

func (r *teacherRepository) UpdateTeacher(teacher models.Teacher) (models.Teacher, error) {
	err := r.db.Save(&teacher).Error
	return teacher, err
}

func (r *teacherRepository) FindByStatus(status string) ([]models.Teacher, error) {
	var teachers []models.Teacher
	err := r.db.Where("status = ?", status).Find(&teachers).Error
	return teachers, err
}

func (r *teacherRepository) DeleteTeacher(nip string) error {

	return r.db.Where("nip = ? ", nip).Delete(&models.Teacher{}).Error
}
