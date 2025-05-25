package usuario

import (
	"net/http"
	"proyecto-integrador/dto"
	"proyecto-integrador/services"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func Login(ctx *gin.Context) {
	var loginJSON dto.UsuarioLoginDTO
	if err := ctx.BindJSON(&loginJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos con formato incorrecto"})
		log.Debug("LoginDTO:", loginJSON)
		return
	}

	token, err := services.UsuarioService.GenerateToken(loginJSON.Username, loginJSON.Password)
	if err != nil {
		log.Debug(err)
		if err == services.IncorrectCredentialsError {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ocurrio un error en el servidor"})
			return
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"access_token": token,
		"token_type":   "bearer",
		"expires_in":   1800, // en segundos
	})
}
