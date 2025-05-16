package usuario

import (
	"proyecto-integrador/db"
	"proyecto-integrador/model"

	log "github.com/sirupsen/logrus"
)

func GetUsuarioByUsername(username string) model.Usuario {
	var usuario model.Usuario
	db.GetInstance().Where("username = ?", username).First(&usuario)

	log.Debug("Usuario: ", usuario)

	return usuario
}
