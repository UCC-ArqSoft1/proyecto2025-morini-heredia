package actividad

import (
	"fmt"
	"proyecto-integrador/db"
	"proyecto-integrador/model"

	log "github.com/sirupsen/logrus"
)

func GetActividadesByParams(params map[string]any) model.Actividades {
	var actividades model.Actividades
	query := db.GetInstance().Model(&model.Actividad{})

	if params["id"] != "" {
		query = query.Where("id_actividad = ?", params["id"])
	}
	if params["titulo"] != "" {
		query = query.Where("titulo LIKE ?", fmt.Sprintf("%%%s%%", params["titulo"]))
	}
	if params["horario"] != "" {
		query = query.Where("TIME(horario) = ?", params["horario"])
	}
	if params["categoria"] != "" {
		query = query.Where("categoria LIKE ?", fmt.Sprintf("%%%s%%", params["categoria"]))
	}

	var err error = query.Find(&actividades).Error
	if err != nil {
		log.Error("Error al buscar actividades: ", err)
	}

	log.Debug("Actividades: ", actividades)

	return actividades
}

func GetActividadById(id int) model.Actividad {
	var actividad model.Actividad
	db.GetInstance().Where("id_actividad = ?", id).First(&actividad)

	log.Debug("Actividad: ", actividad)

	return actividad
}

func GetAllActividades() model.Actividades {
	var actividades model.Actividades
	db.GetInstance().Find(&actividades)

	log.Debug("Actividades: ", actividades)

	return actividades
}
