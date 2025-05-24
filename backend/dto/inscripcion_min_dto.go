package dto

type InscripcionMinDTO struct {
	Id       uint   `json:"id"`
	IsActiva string `json:"is_activa"`
}

type InscripcionesMinDTO []InscripcionMinDTO
