package models

// Enum for Student Status
type StudentStatus string

const (
	Active       StudentStatus = "active"
	StudentLeave StudentStatus = "student leave"
	Out          StudentStatus = "out"
)

type Student struct {
	NIM    string        `json:"nim" gorm:"primaryKey;type:varchar(20);unique;not null"`
	Name   string        `json:"name" binding:"required"`
	Kelas  string        `json:"kelas" binding:"required"`
	Status StudentStatus `json:"status" gorm:"type:varchar(20);not null;default:'active'"`
}
