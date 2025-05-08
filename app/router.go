package app

import (
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

	backend:
	https://localhost:8080/actividades
	https://localhost:8080/actividades/{id}
	https://localhost:8080/usuarios
	https://localhost:8080/usuarios/{id}
*/

func init() {
	router = gin.Default()
	router.Use(cors.Default())
}

func StartRoute() {
	log.Info("Iniciando servidor")
	router.Run(":8080")
}
