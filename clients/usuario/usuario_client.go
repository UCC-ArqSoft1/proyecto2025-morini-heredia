package usuario

import (
	"proyecto-integrador/db"
	"proyecto-integrador/model"

	"gorm.io/gorm"
)

var db_conn *gorm.DB = db.GetInstance()

func GetUsuarioByUsername(username string) (model.Usuario, error) {
	var usuario model.Usuario
	err := db_conn.Where("username = ?", username).First(&usuario).Error

	return usuario, err
}
