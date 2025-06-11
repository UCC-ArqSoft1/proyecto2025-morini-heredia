package services

import (
	"fmt"
	"log"
	"proyecto-integrador/clients/actividad"
	"proyecto-integrador/dto"
	"proyecto-integrador/model"
	"time"
)

type actividadService struct{}

type IactividadService interface {
	GetAllActividades() (dto.ActividadesDTO, error)
	GetActividadesByParams(params map[string]any) (dto.ActividadesDTO, error)
	GetActividadByID(id int) (dto.ActividadDTO, error)
	DeleteActividad(id uint) error
	CreateActividad(actividadDTO dto.ActividadDTO) error
	UpdateActividad(actividadDTO dto.ActividadDTO) error
}

var (
	ActividadService IactividadService
)

func init() {
	ActividadService = &actividadService{}
}

func (s *actividadService) GetAllActividades() (dto.ActividadesDTO, error) {
	var actividades model.ActividadesVista = actividad.GetAllActividades()
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
			Lugares:     v.Lugares,
		}

		actividadesDTO[i] = actividadDTO
	}

	return actividadesDTO, nil
}

func (s *actividadService) GetActividadesByParams(params map[string]any) (dto.ActividadesDTO, error) {
	var actividades model.ActividadesVista = actividad.GetActividadesByParams(params)
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
	var actividad model.ActividadVista = actividad.GetActividadById(id)
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

func (s *actividadService) CreateActividad(actividadDTO dto.ActividadDTO) error {
	log.Printf("Recibiendo DTO para crear actividad: %+v\n", actividadDTO)

	horaInicio, err := time.Parse("15:04", actividadDTO.HoraInicio)
	if err != nil {
		log.Printf("Error al parsear hora inicio '%s': %v\n", actividadDTO.HoraInicio, err)
		return fmt.Errorf("formato de hora inicio inválido: %v", err)
	}

	horaFin, err := time.Parse("15:04", actividadDTO.HoraFin)
	if err != nil {
		log.Printf("Error al parsear hora fin '%s': %v\n", actividadDTO.HoraFin, err)
		return fmt.Errorf("formato de hora fin inválido: %v", err)
	}

	nuevaActividad := model.Actividad{
		Titulo:        actividadDTO.Titulo,
		Descripcion:   actividadDTO.Descripcion,
		Cupo:          actividadDTO.Cupo,
		Dia:           actividadDTO.Dia,
		HorarioInicio: model.CustomTime(horaInicio),
		HorarioFinal:  model.CustomTime(horaFin),
		FotoUrl:       "SAMPLE_URL",
		Instructor:    actividadDTO.Instructor,
		Categoria:     actividadDTO.Categoria,
	}

	log.Printf("Creando actividad: %+v\n", nuevaActividad)
	return actividad.CreateActividad(nuevaActividad)
}

func (s *actividadService) UpdateActividad(actividadDTO dto.ActividadDTO) error {
	log.Printf("Recibiendo DTO para actualizar actividad: %+v\n", actividadDTO)

	horaInicio, err := time.Parse("15:04", actividadDTO.HoraInicio)
	if err != nil {
		log.Printf("Error al parsear hora inicio '%s': %v\n", actividadDTO.HoraInicio, err)
		return fmt.Errorf("formato de hora inicio inválido: %v", err)
	}

	horaFin, err := time.Parse("15:04", actividadDTO.HoraFin)
	if err != nil {
		log.Printf("Error al parsear hora fin '%s': %v\n", actividadDTO.HoraFin, err)
		return fmt.Errorf("formato de hora fin inválido: %v", err)
	}

	actividadActualizada := model.Actividad{
		Id:            actividadDTO.Id,
		Titulo:        actividadDTO.Titulo,
		Descripcion:   actividadDTO.Descripcion,
		Cupo:          actividadDTO.Cupo,
		Dia:           actividadDTO.Dia,
		HorarioInicio: model.CustomTime(horaInicio),
		HorarioFinal:  model.CustomTime(horaFin),
		FotoUrl:       "SAMPLE_URL",
		Instructor:    actividadDTO.Instructor,
		Categoria:     actividadDTO.Categoria,
	}

	log.Printf("Actualizando actividad: %+v\n", actividadActualizada)
	return actividad.UpdateActividad(actividadActualizada)
}

func (s *actividadService) DeleteActividad(id uint) error {
	return actividad.DeleteActividad(id)
}
