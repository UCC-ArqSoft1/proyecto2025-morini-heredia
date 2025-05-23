package services

import (
	"proyecto-integrador/clients/inscripcion"
	"proyecto-integrador/dto"
)

type inscripcionService struct{}

type IInscripcionService interface {
	GetAllInscripciones(id_usuario uint) (dto.InscripcionesDTO, error)
	InscribirUsuario(id_usuario, actividad_id uint) (uint, error)
	DesinscribirUsuario() error
}

var (
	InscripcionService IInscripcionService
)

func init() {
	InscripcionService = &inscripcionService{}
}

func (is *inscripcionService) GetAllInscripciones(id_usuario uint) (dto.InscripcionesDTO, error) {
	inscripciones, err := inscripcion.GetAllInscripciones(id_usuario)
	if err != nil {
		return nil, err
	}

	var resultado dto.InscripcionesDTO
	for _, v := range inscripciones {
		dto := dto.InscripcionDTO{
			Id:                 v.Id,
			Fecha_inscripcion:  v.Fecha_inscripcion,
			Estado_inscripcion: v.Estado_inscripcion,
		}
		resultado = append(resultado, dto)
	}

	return resultado, nil
}

func (is *inscripcionService) InscribirUsuario(id_usuario, actividad_id uint) (uint, error) {
	return 0, nil
}

func (is *inscripcionService) DesinscribirUsuario() error {
	return nil

}
