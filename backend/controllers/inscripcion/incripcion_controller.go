package inscripcion

import (
	"net/http"
	"proyecto-integrador/dto"
	"proyecto-integrador/services"

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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al inscribir al usuario"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": inscId})
}

func ActualizarUsuario(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
