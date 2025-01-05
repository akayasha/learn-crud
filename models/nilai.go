package models

type Nilai struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	NIM       string  `json:"nim" gorm:"type:varchar(20);not null;index"`
	Grade     float64 `json:"grade" gorm:"not null" binding:"required,gte=0,lte=100"`
	Subject   Subject `json:"subject" gorm:"foreignKey:SubjectID"`
	SubjectID uint    `json:"subject_id"`
}
