package dto

import "time"

type InscripcionDTO struct {
	Id               uint      `json:"id"`
	FechaInscripcion time.Time `json:"fecha_inscripcion"`
	IsActiva         bool      `json:"is_activa"`
}

type InscripcionesDTO []InscripcionDTO
