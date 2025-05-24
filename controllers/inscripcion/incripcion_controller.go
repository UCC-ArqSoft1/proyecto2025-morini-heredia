package inscripcion

import (
	"proyecto-integrador/services"

	"github.com/gin-gonic/gin"
)

// TODO: esto es para la entrega final
func GetAllInscripciones(ctx *gin.Context) {
	// userID, exists := ctx.Get("user_id")
	// if !exists {
	// 	ctx.JSON(401, gin.H{"error": "no autorizado"})
	// }

	inscripciones, err := services.InscripcionService.GetAllInscripciones(2)
	if err != nil {
		// errores van según si el usuario existe o no, por ahora solo dejo este
		ctx.JSON(500, gin.H{"error": "error al procesar la consulta"})
	}

	ctx.JSON(200, inscripciones)
}

func InscribirUsuario(ctx *gin.Context) {
	// userID, _ := ctx.Get("user_id")

	// var actividad_id dto.IdDTO
	// if err := ctx.BindJSON(&actividad_id); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos con formato incorrecto"})
	// 	log.Debug("IdDTO:", actividad_id)
	// 	return
	// }

	// err := inscr

	ctx.JSON(201, gin.H{"id": "420"})
}

// TODO: terminar de implementar
func ActualizarUsuario(ctx *gin.Context) {
	ctx.Status(204)
}
