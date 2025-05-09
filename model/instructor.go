package model

type Instructor struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Nombre   string `gorm:"type:varchar(100);not null"`
	Apellido string `gorm:"type:varchar(100);not null"`
}

type Instructores []Instructor
