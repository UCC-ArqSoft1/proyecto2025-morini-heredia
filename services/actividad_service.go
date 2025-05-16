package services

import (
	"fmt"
	"proyecto-integrador/clients/actividad"
	"proyecto-integrador/dto"
	"proyecto-integrador/model"
)

type actividadService struct{}

type actividadServiceInterface interface {
	GetAllActividades() (dto.ActividadesMinDTO, error)
	GetActividadesByParams(params map[string]any) (dto.ActividadesMinDTO, error)
	GetActividadByID(id int) (dto.ActividadDTO, error)
}

var (
	ActividadService actividadServiceInterface
)

func init() {
	ActividadService = &actividadService{}
}

func (s *actividadService) GetAllActividades() (dto.ActividadesMinDTO, error) {
	var actividades model.Actividades = actividad.GetAllActividades()
	var actividadesDTO dto.ActividadesMinDTO = make(dto.ActividadesMinDTO, len(actividades))

	for i, v := range actividades {
		actividadDTO := dto.ActividadMinDTO{
			Titulo:      v.Titulo,
			Descripcion: v.Descripcion,
			Cupo:        v.Cupo,
			Dia:         v.Dia,
			HoraInicio:  v.Horario_inicio.Format("15:04"),
			HoraFin:     v.Horario_final.Format("15:04"),
			Instructor:  v.Instructor,
			Categoria:   v.Categoria,
		}

		actividadesDTO[i] = actividadDTO
	}

	return actividadesDTO, nil
}

func (s *actividadService) GetActividadesByParams(params map[string]any) (dto.ActividadesMinDTO, error) {
	var actividades model.Actividades = actividad.GetActividadesByParams(params)
	var actividadesDTO dto.ActividadesMinDTO = make(dto.ActividadesMinDTO, len(actividades))

	for i, v := range actividades {
		actividadDTO := dto.ActividadMinDTO{
			Titulo:      v.Titulo,
			Descripcion: v.Descripcion,
			Cupo:        v.Cupo,
			Dia:         v.Dia,
			HoraInicio:  v.Horario_inicio.Format("15:04"),
			HoraFin:     v.Horario_final.Format("15:04"),
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
		HoraInicio:  actividad.Horario_inicio.Format("15:04"),
		HoraFin:     actividad.Horario_final.Format("15:04"),
		Instructor:  actividad.Instructor,
		Categoria:   actividad.Categoria,
	}

	return actividadDTO, nil
}
