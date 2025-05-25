package inscripcion

import (
	"proyecto-integrador/db"
	"proyecto-integrador/model"

	log "github.com/sirupsen/logrus"
)

func GetAllInscripciones(id_usuario uint) (model.Inscripciones, error) {
	var inscripciones model.Inscripciones
	query := db.GetInstance().Model(&model.Inscripcion{})

	var err error
	if err = query.Where("id_usuario = ?", id_usuario).Find(&inscripciones).Error; err != nil {
		log.Error("Error al buscar inscripciones: ", err)
		return nil, err
	}

	log.Debug("Inscripciones: ", inscripciones)

	return inscripciones, nil
}

func InsertarInscripcion(id_usuario, id_actividad uint) (uint, error) {
	insc := model.Inscripcion{
		IdUsuario:   id_usuario,
		IdActividad: id_actividad,
	}

	return id_usuario, db.GetInstance().Create(&insc).Error
}

// func DesactivarInscripcion(insc model.Inscripcion) model.Inscripcion {

// }
