package services

import (
	"fmt"
	"proyecto-integrador/clients/actividad"
	"proyecto-integrador/dto"
	"proyecto-integrador/model"
)

type actividadService struct{}

type IactividadService interface {
	GetAllActividades() (dto.ActividadesDTO, error)
	GetActividadesByParams(params map[string]any) (dto.ActividadesDTO, error)
	GetActividadByID(id int) (dto.ActividadDTO, error)
	DeleteActividad(id uint) error
}

var (
	ActividadService IactividadService
)

func init() {
	ActividadService = &actividadService{}
}

func (s *actividadService) GetAllActividades() (dto.ActividadesDTO, error) {
	var actividades model.Actividades = actividad.GetAllActividades()
	var actividadesDTO dto.ActividadesDTO = make(dto.ActividadesDTO, len(actividades))

	for i, v := range actividades {
		actividadDTO := dto.ActividadDTO{
			Id:          v.Id,
			Titulo:      v.Titulo,
			Descripcion: v.Descripcion,
			Cupo:        v.Cupo,
			Dia:         v.Dia,
			HoraInicio:  v.HorarioInicio.Format("15:04"),
			HoraFin:     v.HorarioFinal.Format("15:04"),
			Instructor:  v.Instructor,
			Categoria:   v.Categoria,
		}

		actividadesDTO[i] = actividadDTO
	}

	return actividadesDTO, nil
}

func (s *actividadService) GetActividadesByParams(params map[string]any) (dto.ActividadesDTO, error) {
	var actividades model.Actividades = actividad.GetActividadesByParams(params)
	var actividadesDTO dto.ActividadesDTO = make(dto.ActividadesDTO, len(actividades))

	for i, v := range actividades {
		actividadDTO := dto.ActividadDTO{
			Id:          v.Id,
			Titulo:      v.Titulo,
			Descripcion: v.Descripcion,
			Cupo:        v.Cupo,
			Dia:         v.Dia,
			HoraInicio:  v.HorarioInicio.Format("15:04"),
			HoraFin:     v.HorarioFinal.Format("15:04"),
			Instructor:  v.Instructor,
			Categoria:   v.Categoria,
		}

		actividadesDTO[i] = actividadDTO
	}

	return actividadesDTO, nil
}

func (s *actividadService) GetActividadByID(id int) (dto.ActividadDTO, error) {
	var actividad model.Actividad = actividad.GetActividadById(id)
	if actividad.Id == 0 {
		return dto.ActividadDTO{}, fmt.Errorf("actividad con ID %d no encontrada", id)
	}

	var actividadDTO dto.ActividadDTO = dto.ActividadDTO{
		Id:          actividad.Id,
		Titulo:      actividad.Titulo,
		Descripcion: actividad.Descripcion,
		Cupo:        actividad.Cupo,
		Dia:         actividad.Dia,
		HoraInicio:  actividad.HorarioInicio.Format("15:04"),
		HoraFin:     actividad.HorarioFinal.Format("15:04"),
		Instructor:  actividad.Instructor,
		Categoria:   actividad.Categoria,
	}

	return actividadDTO, nil
}

func (s *actividadService) DeleteActividad(id uint) error {
	return actividad.DeleteActividad(id)
}
