package app

import (
	"net/http"
	"os"
	"proyecto-integrador/clients/actividad"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

/*
	frontend:
	https://localhost:5001/home
	https://localhost:5001/login
	https://localhost:5001/signup
	https://localhost:5001/activities
	https://localhost:5001/user/activities -- estar loggeado

	backend:
	https://localhost:8080/actividades?{id,titulo,horario,categoria}
	https://localhost:8080/usuarios?{id}
	https://localhost:8080/usuarios/actividades
*/

func init() {
	router = gin.Default()
	router.Use(cors.Default())
}

func BuscarActividades(ctx *gin.Context) {
	titulo := ctx.Query("titulo")
	horario := ctx.Query("horario")
	categoria := ctx.Query("categoria")

	actividades, err := actividad.GetActividades(map[string]any{"titulo": titulo, "horario": horario, "categoria": categoria})
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error al buscar actividades"})
	}

	ctx.JSON(200, actividades)
}

func getActividades(ctx *gin.Context) {
	actividades := actividad.GetAllActividades()

	ctx.JSON(http.StatusOK, actividades)
}

func getActividadById(ctx *gin.Context) {
	id_actividad, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "el id debe ser un numero"})
		return
	}

	actividad := actividad.GetActividadById(id_actividad)

	ctx.JSON(http.StatusOK, actividad)
}

// TODO: terminar la implementación de estas funciones
func getActividad(ctx *gin.Context) {

}

func getActividadesUsuario(ctx *gin.Context) {

}

func login(ctx *gin.Context) {
	// va
}

func signup(ctx *gin.Context) {

}

func StartRoute() {
	// TODO: declarar endpoints de la app

	router.GET("/actividades", getActividades)
	router.GET("/actividades/:id", getActividadById)
	router.GET("/usuario/actividades", getActividadesUsuario)

	router.POST("/login", login)
	router.POST("/signup", signup)

	host := os.Getenv("APP_HOST")
	if host == "" {
		host = ":8080"
	}

	log.Info("Iniciando servidor en: ", host)
	router.Run(host)
}
