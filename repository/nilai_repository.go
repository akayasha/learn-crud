package repository

import (
	"gorm.io/gorm"
	"learn-crud/models"
)

type NilaiRepository interface {
	FindByNIM(nim string) ([]models.Nilai, error)
	Create(nilai models.Nilai) (models.Nilai, error)
	Update(nilai models.Nilai) (models.Nilai, error)
	Delete(id uint) error
}

type nilaiRepository struct {
	db *gorm.DB
}

func NewNilaiRepository(db *gorm.DB) NilaiRepository {
	return &nilaiRepository{db: db}
}

func (r *nilaiRepository) FindByNIM(nim string) ([]models.Nilai, error) {
	var nilai []models.Nilai
	// Fetching all grades associated with a student's NIM
	err := r.db.Where("nim = ?", nim).Find(&nilai).Error
	return nilai, err
}

func (r *nilaiRepository) Create(nilai models.Nilai) (models.Nilai, error) {
	// Creating a new grade for a student
	err := r.db.Create(&nilai).Error
	return nilai, err
}

func (r *nilaiRepository) Update(nilai models.Nilai) (models.Nilai, error) {
	err := r.db.Save(&nilai).Error
	return nilai, err
}

func (r *nilaiRepository) Delete(id uint) error {
	// Deleting a grade by ID
	return r.db.Delete(&models.Nilai{}, id).Error
}
