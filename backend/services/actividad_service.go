package services

import (
	"errors"
	"fmt"
	"proyecto-integrador/clients/actividad"
	"proyecto-integrador/dto"
	"proyecto-integrador/model"
	"time"

	log "github.com/sirupsen/logrus"
)

type actividadService struct{}

type IactividadService interface {
	GetAllActividades() (dto.ActividadesMinDTO, error)
	GetActividadesByParams(params map[string]any) (dto.ActividadesMinDTO, error)
	GetActividadByID(id int) (dto.ActividadDTO, error)
	DeleteActividad(id uint) error
	UpdateActividad(id uint, actividadDTO dto.ActividadDTO) error
}

var (
	ActividadService IactividadService
)

func init() {
	ActividadService = &actividadService{}
}

func (s *actividadService) GetAllActividades() (dto.ActividadesMinDTO, error) {
	var actividades model.Actividades = actividad.GetAllActividades()
	var actividadesDTO dto.ActividadesMinDTO = make(dto.ActividadesMinDTO, len(actividades))

	for i, v := range actividades {
		actividadDTO := dto.ActividadMinDTO{
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

func (s *actividadService) GetActividadesByParams(params map[string]any) (dto.ActividadesMinDTO, error) {
	var actividades model.Actividades = actividad.GetActividadesByParams(params)
	var actividadesDTO dto.ActividadesMinDTO = make(dto.ActividadesMinDTO, len(actividades))

	for i, v := range actividades {
		actividadDTO := dto.ActividadMinDTO{
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

func (s *actividadService) UpdateActividad(id uint, actividadDTO dto.ActividadDTO) error {
	// Validar que la actividad exista
	existingActividad := actividad.GetActividadById(int(id))
	if existingActividad.Id == 0 {
		return fmt.Errorf("actividad con ID %d no encontrada", id)
	}

	// Validar los datos recibidos
	if actividadDTO.Titulo == "" {
		return errors.New("el título no puede estar vacío")
	}
	if actividadDTO.Cupo <= 0 {
		return errors.New("el cupo debe ser mayor a 0")
	}
	if actividadDTO.Dia == "" {
		return errors.New("el día no puede estar vacío")
	}
	if actividadDTO.HoraInicio == "" || actividadDTO.HoraFin == "" {
		return errors.New("las horas de inicio y fin son requeridas")
	}
	if actividadDTO.Instructor == "" {
		return errors.New("el instructor no puede estar vacío")
	}
	if actividadDTO.Categoria == "" {
		return errors.New("la categoría no puede estar vacía")
	}

	// Convertir el DTO a modelo
	horaInicio, err := time.Parse("15:04", actividadDTO.HoraInicio)
	if err != nil {
		return fmt.Errorf("formato de hora de inicio inválido: %v", err)
	}

	horaFin, err := time.Parse("15:04", actividadDTO.HoraFin)
	if err != nil {
		return fmt.Errorf("formato de hora de fin inválido: %v", err)
	}

	// Validar que la hora de fin sea posterior a la hora de inicio
	if horaFin.Before(horaInicio) {
		return errors.New("la hora de fin debe ser posterior a la hora de inicio")
	}

	// Usar la fecha actual como base para las horas
	now := time.Now()
	horaInicio = time.Date(now.Year(), now.Month(), now.Day(), horaInicio.Hour(), horaInicio.Minute(), 0, 0, time.Local)
	horaFin = time.Date(now.Year(), now.Month(), now.Day(), horaFin.Hour(), horaFin.Minute(), 0, 0, time.Local)

	actividadModel := model.Actividad{
		Id:            id,
		Titulo:        actividadDTO.Titulo,
		Descripcion:   actividadDTO.Descripcion,
		Cupo:          actividadDTO.Cupo,
		Dia:           actividadDTO.Dia,
		HorarioInicio: horaInicio,
		HorarioFinal:  horaFin,
		Instructor:    actividadDTO.Instructor,
		Categoria:     actividadDTO.Categoria,
		FotoUrl:       existingActividad.FotoUrl, // Mantener la URL de foto existente
	}

	log.Info("Actualizando actividad con ID:", id)     // Para depuración
	log.Info("Datos de la actividad:", actividadModel) // Para depuración

	return actividad.UpdateActividad(actividadModel)
}
