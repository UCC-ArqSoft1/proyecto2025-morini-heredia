package inscripcion

import (
	"net/http"
	"proyecto-integrador/dto"
	"proyecto-integrador/services"
	"strings"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func GetAllInscripciones(ctx *gin.Context) {
	userID, exists := ctx.Get("id_usuario")
	if !exists {
		log.Error("la variable 'id_usuario' no esta definida")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	inscripciones, err := services.InscripcionService.GetAllInscripciones(userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error al procesar la consulta"})
		return
	}

	ctx.JSON(http.StatusOK, inscripciones)
}

func InscribirUsuario(ctx *gin.Context) {
	userID, exists := ctx.Get("id_usuario")
	if !exists {
		log.Error("la variable 'id_usuario' no esta definida")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	var idDTO dto.IdDTO
	if err := ctx.BindJSON(&idDTO); err != nil {
		log.Debug("IdDTO:", idDTO)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos con formato incorrecto"})
		return
	}

	inscId, err := services.InscripcionService.InscribirUsuario(userID.(uint), idDTO.Id)
	if err != nil {
		log.Error(err)

		// TODO: ver si se puede revisar de otra forma el error que no sea mediante strings
		errString := strings.ToLower(err.Error())
		if strings.Contains(errString, "error 1062") {
			// if errors.Is(err, gorm.ErrDuplicatedKey) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "El usuario ya esta inscripto a esta actividad"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al inscribir al usuario"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": inscId})
}

// TODO: implementar
func DesinscribirUsuario(ctx *gin.Context) {
	userID, exists := ctx.Get("id_usuario")
	if !exists {
		log.Error("la variable 'id_usuario' no esta definida")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	var idDTO dto.IdDTO
	if err := ctx.BindJSON(&idDTO); err != nil {
		log.Debug("IdDTO:", idDTO)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos con formato incorrecto"})
		return
	}

	err := services.InscripcionService.DesinscribirUsuario(userID.(uint), idDTO.Id)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al inscribir al usuario"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
