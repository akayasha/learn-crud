package models

type Subject struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name" gorm:"type:varchar(100);not null"`
	Teacher   Teacher `json:"teacher" gorm:"foreignKey:TeacherID"`
	TeacherID uint    `json:"teacher_id"`
}
