package inscripcion

import (
	"proyecto-integrador/db"
	"proyecto-integrador/model"

	"gorm.io/gorm"
)

var db_conn *gorm.DB = db.GetInstance()

// TODO: terminar de implementar esto
func InsertarInscripcion(insc model.Inscripcion) model.Inscripcion

func BorrarInscripcion(insc model.Inscripcion) model.Inscripcion
