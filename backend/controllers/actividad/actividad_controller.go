package actividad

import (
	"net/http"
	"proyecto-integrador/dto"
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
		ctx.JSON(404, gin.H{"error": "La actividad no existe"})
		return
	}

	ctx.JSON(http.StatusOK, actividad)
}

func DeleteActividad(ctx *gin.Context) {
	isAdmin, exists := ctx.Get("is_admin")
	if !exists || !isAdmin.(bool) {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos para realizar esta acción"})
		return
	}

	idActividad, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un número"})
		return
	}

	err = services.ActividadService.DeleteActividad(uint(idActividad))
	if err != nil {
		log.Error("Error al eliminar actividad:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la actividad"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func UpdateActividad(ctx *gin.Context) {
	log.Info("Recibida solicitud PUT /actividades/:id")

	// Verificar si el usuario es admin
	isAdmin, exists := ctx.Get("is_admin")
	if !exists || !isAdmin.(bool) {
		log.Error("Usuario no autorizado intentando actualizar actividad")
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": "error",
			"error":  "No tienes permisos para realizar esta acción",
		})
		return
	}

	// Obtener el ID de la actividad
	idActividad, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Error("Error al convertir ID:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "El id debe ser un número",
		})
		return
	}

	log.Infof("Actualizando actividad con ID: %d", idActividad)

	// Obtener los datos actualizados del body
	var actividadDTO dto.ActividadDTO
	if err := ctx.ShouldBindJSON(&actividadDTO); err != nil {
		log.Error("Error al parsear datos:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "error",
			"error":    "Datos inválidos",
			"detalles": err.Error(),
		})
		return
	}

	log.Info("Datos recibidos:", actividadDTO)

	// Llamar al servicio para actualizar la actividad
	err = services.ActividadService.UpdateActividad(uint(idActividad), actividadDTO)
	if err != nil {
		log.Error("Error al actualizar actividad:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	log.Info("Actividad actualizada correctamente")
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"mensaje": "Actividad actualizada correctamente",
	})
}
