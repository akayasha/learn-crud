package models

// Enum for status
type TeacherStatus string

const (
	TeacherActive TeacherStatus = "active"
	Inactive      TeacherStatus = "inactive"
)

type Teacher struct {
	NIP    string        `json:"nip" gorm:"primaryKey;type:varchar(20);unique;not null"`
	Name   string        `json:"name" gorm:"type:varchar(100);not null"`
	Status TeacherStatus `json:"status" gorm:"type:varchar(100);not null"`
}
