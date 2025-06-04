package model

import "time"

type Inscripcion struct {
	IdUsuario        uint      `gorm:"column:id_usuario;primaryKey"`
	IdActividad      uint      `gorm:"column:id_actividad;primaryKey"`
	FechaInscripcion time.Time `gorm:"column:fecha_inscripcion;type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	IsActiva         bool      `gorm:"column:is_activa;default:true;not null"`

	Usuario   Usuario   `gorm:"foreignKey:IdUsuario"`
	Actividad Actividad `gorm:"foreignKey:IdActividad"`
}

type Inscripciones []Inscripcion
