package inscripcion

import (
	"proyecto-integrador/db"
	"proyecto-integrador/model"

	log "github.com/sirupsen/logrus"
)

func GetAllInscripciones(usuario_id uint) (model.Inscripciones, error) {
	var inscripciones model.Inscripciones
	query := db.GetInstance().Model(&model.Inscripcion{})

	var err error
	if err = query.Where("usuario_id = ?", usuario_id).Find(&inscripciones).Error; err != nil {
		log.Error("Error al buscar inscripciones: ", err)
		return nil, err
	}

	log.Debug("Inscripciones: ", inscripciones)

	return inscripciones, nil
}

// func InsertarInscripcion(usuario_id, actividad_id uint) (model.Inscripcion, error) {
// 	inscripcion := model.Inscripcion{
// 		Usuario_id:         usuario_id,
// 		Actividad_id:       actividad_id,
// 		Estado_inscripcion: "activa",
// 	}

// 	db.GetInstance().
// 	return inscripcion, nil
// }

// func DesactivarInscripcion(insc model.Inscripcion) model.Inscripcion {

// }
