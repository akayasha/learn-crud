package services

import (
	"learn-crud/models"
	"learn-crud/repository"
)

type TeacherService interface {
	GetAllTeachers() ([]models.Teacher, error)
	GetTeacherByNIP(nip string) (models.Teacher, error)
	CreateTeacher(teacher models.Teacher) (models.Teacher, error)
	UpdateTeacher(teacher models.Teacher) (models.Teacher, error)
	DeleteTeacher(nip string) error
	GetTeachersByStatus(status string) ([]models.Teacher, error)
}

type teacherService struct {
	repo repository.TeacherRepository
}

func NewTeacherService(repo repository.TeacherRepository) TeacherService {
	return &teacherService{repo: repo}
}

func (s *teacherService) GetAllTeachers() ([]models.Teacher, error) {
	return s.repo.FindAll()
}

func (s *teacherService) GetTeacherByNIP(nip string) (models.Teacher, error) {
	return s.repo.FindByNip(nip)
}

func (s *teacherService) CreateTeacher(teacher models.Teacher) (models.Teacher, error) {
	return s.repo.CreateTeacher(teacher)
}

func (s *teacherService) UpdateTeacher(teacher models.Teacher) (models.Teacher, error) {
	return s.repo.UpdateTeacher(teacher)
}

func (s *teacherService) DeleteTeacher(nip string) error {
	return s.repo.DeleteTeacher(nip)
}

func (s *teacherService) GetTeachersByStatus(status string) ([]models.Teacher, error) {
	return s.repo.FindByStatus(status)
}
