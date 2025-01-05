package services

import (
	"learn-crud/models"
	"learn-crud/repository"
)

type NilaiService interface {
	GetNilaiByNIM(nim string) ([]models.Nilai, error)
	CreateNilai(nilai models.Nilai) (models.Nilai, error)
	UpdateNilai(nilai models.Nilai) (models.Nilai, error)
	DeleteNilai(id uint) error
}

type nilaiService struct {
	repo repository.NilaiRepository
}

func NewNilaiService(repo repository.NilaiRepository) NilaiService {
	return &nilaiService{repo: repo}
}

func (s *nilaiService) GetNilaiByNIM(nim string) ([]models.Nilai, error) {
	return s.repo.FindByNIM(nim)
}

func (s *nilaiService) CreateNilai(nilai models.Nilai) (models.Nilai, error) {
	return s.repo.Create(nilai)
}

func (s *nilaiService) UpdateNilai(nilai models.Nilai) (models.Nilai, error) {
	return s.repo.Update(nilai)
}

func (s *nilaiService) DeleteNilai(id uint) error {
	return s.repo.Delete(id)
}
