package model

import "time"

type Inscripcion struct {
	Id                 uint      `gorm:"primaryKey;autoIncrement"`
	Fecha_inscripcion  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	Estado_inscripcion string    `gorm:"type:enum('activa','inactiva');not null;default:'activa'"`

	Usuario_iD   uint `gorm:"not null"` // FK a Usuario
	Actividad_iD uint `gorm:"not null"` // FK a Actividad

	Usuario   Usuario   `gorm:"foreignKey:UsuarioId"`
	Actividad Actividad `gorm:"foreignKey:ActividadId"`
}

type Inscripciones []Inscripcion
