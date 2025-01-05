package repository

import (
	"gorm.io/gorm"
	"learn-crud/models"
)

type StudentRepository interface {
	FindAll() ([]models.Student, error)
	FindByNIM(nim string) (models.Student, error)
	Create(student models.Student) (models.Student, error)
	Update(student models.Student) (models.Student, error)
	Delete(nim string) error
	FindByKelas(kelas string) ([]models.Student, error)
	FindByName(name string) ([]models.Student, error)
	FindByStatus(status string) ([]models.Student, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) FindByStatus(status string) ([]models.Student, error) {
	var students []models.Student
	err := r.db.Where("status = ?", status).Find(&students).Error
	return students, err
}
func (r *studentRepository) FindByKelas(kelas string) ([]models.Student, error) {
	var students []models.Student
	// Fetching students by class (kelas), and it will return multiple students matching the class.
	err := r.db.Where("kelas = ?", kelas).Find(&students).Error
	return students, err
}

func (r *studentRepository) FindByName(name string) ([]models.Student, error) {
	var students []models.Student
	// Fetching students whose name contains the search term (case-insensitive).
	err := r.db.Where("name LIKE ?", "%"+name+"%").Find(&students).Error
	return students, err
}

func (r *studentRepository) FindAll() ([]models.Student, error) {
	var students []models.Student
	// Fetching all students.
	err := r.db.Find(&students).Error
	return students, err
}

func (r *studentRepository) FindByNIM(nim string) (models.Student, error) {
	var student models.Student
	// Fetching a single student by NIM (primary key).
	err := r.db.Where("nim = ?", nim).First(&student).Error
	return student, err
}

func (r *studentRepository) Create(student models.Student) (models.Student, error) {
	// Creating a new student in the database.
	err := r.db.Create(&student).Error
	return student, err
}

func (r *studentRepository) Update(student models.Student) (models.Student, error) {
	// Updating an existing student in the database.
	err := r.db.Save(&student).Error
	return student, err
}

func (r *studentRepository) Delete(nim string) error {
	// Deleting a student by NIM.
	return r.db.Where("nim = ?", nim).Delete(&models.Student{}).Error
}
