package app

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartRoute() {
	// Configurar el modo de Gin
	gin.SetMode(gin.DebugMode)

	MapMiddewares()
	MapURLs()

	// Imprimir todas las rutas registradas
	routes := router.Routes()
	log.Info("Rutas registradas:")
	for _, route := range routes {
		log.Infof("%s %s", route.Method, route.Path)
	}

	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	hostport := host + ":" + port

	log.Info("Iniciando servidor en: ", hostport)
	if err := router.Run(hostport); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
