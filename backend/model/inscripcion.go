package model

import "time"

type Inscripcion struct {
	Id               uint      `gorm:"column:id_inscripcion;primaryKey;autoIncrement"`
	FechaInscripcion time.Time `gorm:"column:fecha_inscripcion;type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	IsActiva         bool      `gorm:"column:is_activa;default:true;not null"`

	IdUsuario uint
	Usuario   Usuario `gorm:"foreignKey:IdUsuario"`

	IdActividad uint
	Actividad   Actividad `gorm:"foreignKey:IdActividad"`
}

type Inscripciones []Inscripcion
