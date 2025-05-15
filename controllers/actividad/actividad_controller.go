package actividad

import (
	"net/http"
	"proyecto-integrador/services"
	"strconv"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func GetActividadesByParams(ctx *gin.Context) {
	actividades, err := services.ActividadService.GetActividadesByParams(map[string]any{
		"id":        ctx.Query("id"),
		"titulo":    ctx.Query("titulo"),
		"horario":   ctx.Query("horario"),
		"categoria": ctx.Query("categoria")},
	)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error al buscar actividades"})
	}

	ctx.JSON(200, actividades)
}

func GetAllActividades(ctx *gin.Context) {
	actividades, err := services.ActividadService.GetAllActividades()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar actividades"})
		return
	}

	ctx.JSON(http.StatusOK, actividades)
}

func GetActividadById(ctx *gin.Context) {
	id_actividad, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un numero"})
		return
	}

	actividad, err := services.ActividadService.GetActividadByID(id_actividad)
	if err != nil {
		log.Error("Error al buscar actividad:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "La actividad no existe"})
		return
	}

	ctx.JSON(http.StatusOK, actividad)
}
