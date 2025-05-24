package dto

// TODO: preguntar si está bien hacer un dto solo para enviar esto
type IdDTO struct {
	Id uint `json:"id" binding:"required"`
}
