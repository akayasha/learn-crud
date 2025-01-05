package repository

import (
	"gorm.io/gorm"
	"learn-crud/models"
)

type SubjectRepository interface {
	FindAll() ([]models.Subject, error)
	FindByID(id uint) (models.Subject, error)
	Create(subject models.Subject) (models.Subject, error)
	Update(subject models.Subject) (models.Subject, error)
	Delete(id uint) error
}

type subjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &subjectRepository{db: db}
}

func (r *subjectRepository) FindAll() ([]models.Subject, error) {
	var subjects []models.Subject
	err := r.db.Preload("Teacher").Find(&subjects).Error
	return subjects, err
}

func (r *subjectRepository) FindByID(id uint) (models.Subject, error) {
	var subject models.Subject
	err := r.db.Preload("Teacher").First(&subject, id).Error
	return subject, err
}

func (r *subjectRepository) Create(subject models.Subject) (models.Subject, error) {
	err := r.db.Create(&subject).Error
	return subject, err
}

func (r *subjectRepository) Update(subject models.Subject) (models.Subject, error) {
	err := r.db.Save(&subject).Error
	return subject, err
}

func (r *subjectRepository) Delete(id uint) error {
	return r.db.Delete(&models.Subject{}, id).Error
}
