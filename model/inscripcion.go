package model

import "time"

type Inscripcion struct {
	ID     uint      `gorm:"primaryKey;autoIncrement"`
	Fecha  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	Estado string    `gorm:"type:enum('activa','inactiva');not null;default:'activa'"`

	UsuarioID   uint `gorm:"not null"` // FK a Usuario
	ActividadID uint `gorm:"not null"` // FK a Actividad

	Usuario   Usuario   `gorm:"foreignKey:UsuarioId"`
	Actividad Actividad `gorm:"foreignKey:ActividadId"`
}

type Inscripciones []Inscripcion
