package model

import "time"

type Inscripcion struct {
	Id               uint      `gorm:"column:id_inscripion;primaryKey;autoIncrement"`
	FechaInscripcion time.Time `gorm:"column:fecha_inscripcion;type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	IsActiva         bool      `gorm:"column:is_activa;default:true;not null"`

	UsuarioId uint
	Usuario   Usuario `gorm:"foreignKey:UsuarioId"`

	ActividadId uint
	Actividad   Actividad `gorm:"foreignKey:ActividadId"`
}

type Inscripciones []Inscripcion
