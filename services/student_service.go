package services

import (
	"fmt"
	"learn-crud/models"
	"learn-crud/repository"
)

type StudentService interface {
	GetAllStudents() ([]models.Student, error)
	GetStudentByNIM(nim string) (models.Student, error)
	CreateStudent(student models.Student) (models.Student, error)
	UpdateStudent(student models.Student) (models.Student, error)
	DeleteStudent(nim string) error
	GetStudentByKelas(kelas string) ([]models.Student, error)
	GetStudentByName(name string) ([]models.Student, error)
	GetStudentByStatus(status string) ([]models.Student, error)
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

func (s *studentService) GetAllStudents() ([]models.Student, error) {
	return s.repo.FindAll()
}

func (s *studentService) GetStudentByNIM(nim string) (models.Student, error) {
	return s.repo.FindByNIM(nim)
}

func (s *studentService) GetStudentByStatus(status string) ([]models.Student, error) {
	return s.repo.FindByStatus(status)

}

func (s *studentService) GetStudentByKelas(kelas string) ([]models.Student, error) {
	// This will return a list of students in a specific class (kelas).
	return s.repo.FindByKelas(kelas)
}

func (s *studentService) GetStudentByName(name string) ([]models.Student, error) {
	// This will return a list of students whose name contains the search term.
	return s.repo.FindByName(name)
}

func (s *studentService) CreateStudent(student models.Student) (models.Student, error) {
	_, err := s.repo.Create(student)
	if err != nil {
		// Check for duplicate key error (NIM must be unique).
		if isDuplicateKeyError(err) {
			return models.Student{}, fmt.Errorf("duplicate NIM: %s", student.NIM)
		}
		return models.Student{}, err
	}
	return student, nil
}

func (s *studentService) UpdateStudent(student models.Student) (models.Student, error) {

	if student.Status == "" {
		student.Status = models.Active
	}

	return s.repo.Update(student)
}

func (s *studentService) DeleteStudent(nim string) error {
	return s.repo.Delete(nim)
}

// Helper function to check if the error is a duplicate key error
func isDuplicateKeyError(err error) bool {
	// GORM returns error messages containing "Duplicate entry" for duplicate key errors
	return err != nil && (err.Error() == "Duplicate entry" || err.Error() == "ERROR 1062 (23000): Duplicate entry")
}
