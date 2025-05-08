package actividad

import (
	"fmt"
	"proyecto-integrador/db"
	"proyecto-integrador/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db_conn *gorm.DB = db.GetInstance()

func GetActividades(params map[string]any) (model.Actividades, error) {
	var actividades model.Actividades

	query := db_conn.Model(&model.Actividad{})

	if params["titulo"] != "" {
		query = query.Where("titulo LIKE ?", fmt.Sprintf("%%%s%%", params["titulo"]))
	}
	if params["horario"] != "" {
		query = query.Where("TIME(horario) = ?", params["horario"])
	}
	if params["categoria"] != "" {
		query = query.Where("categoria LIKE ?", fmt.Sprintf("%%%s%%", params["categoria"]))
	}

	err := query.Find(&actividades).Error
	return actividades, err
}

func GetActividadById(id int) model.Actividad {
	var actividad model.Actividad
	db_conn.Where("id = ?", id).First(&actividad)

	log.Debug("Actividad: ", actividad)

	return actividad
}

func GetAllActividades() model.Actividades {
	var actividades model.Actividades
	db_conn.Find(&actividades)

	log.Debug("Actividades: ", actividades)

	return actividades
}
