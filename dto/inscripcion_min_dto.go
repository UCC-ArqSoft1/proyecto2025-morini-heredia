package dto

type InscripcionMinDTO struct {
	Id                 uint   `json:"id"`
	Estado_inscripcion string `json:"estado_inscripcion"`
}

type InscripcionesMinDTO []InscripcionMinDTO
