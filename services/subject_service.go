package services

import (
	"learn-crud/models"
	"learn-crud/repository"
)

type SubjectService interface {
	GetAllSubjects() ([]models.Subject, error)
	GetSubjectByID(id uint) (models.Subject, error)
	CreateSubject(subject models.Subject) (models.Subject, error)
	UpdateSubject(subject models.Subject) (models.Subject, error)
	DeleteSubject(id uint) error
}

type subjectService struct {
	repo repository.SubjectRepository
}

func NewSubjectService(repo repository.SubjectRepository) SubjectService {
	return &subjectService{repo: repo}
}

func (s *subjectService) GetAllSubjects() ([]models.Subject, error) {
	return s.repo.FindAll()
}

func (s *subjectService) GetSubjectByID(id uint) (models.Subject, error) {
	return s.repo.FindByID(id)
}

func (s *subjectService) CreateSubject(subject models.Subject) (models.Subject, error) {
	return s.repo.Create(subject)
}

func (s *subjectService) UpdateSubject(subject models.Subject) (models.Subject, error) {
	return s.repo.Update(subject)
}

func (s *subjectService) DeleteSubject(id uint) error {
	return s.repo.Delete(id)
}
