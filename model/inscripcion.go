package model

import "time"

type Inscripcion struct {
	Id                 uint      `gorm:"primaryKey;autoIncrement"`
	Fecha_inscripcion  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	Estado_inscripcion string    `gorm:"type:enum('activa','inactiva');not null;default:'activa'"`

	Usuario_id   uint `gorm:"not null"` // FK a Usuario
	Actividad_id uint `gorm:"not null"` // FK a Actividad

	Usuario   Usuario   `gorm:"foreignKey:usuario_id"`
	Actividad Actividad `gorm:"foreignKey:actividad_id"`
}

type Inscripciones []Inscripcion
