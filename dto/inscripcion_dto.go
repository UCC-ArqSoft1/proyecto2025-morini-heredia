package dto

import "time"

type InscripcionDTO struct {
	Id                 uint      `json:"id"`
	Fecha_inscripcion  time.Time `json:"fecha_inscripcion"`
	Estado_inscripcion string    `json:"estado_inscripcion"`
}

type InscripcionesDTO []InscripcionDTO
