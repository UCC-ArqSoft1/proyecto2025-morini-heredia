package app

import (
	"net/http"
	"os"
	"proyecto-integrador/clients/usuario"
	"proyecto-integrador/controllers/actividad"
	"proyecto-integrador/dto"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

/*
backend:

GET https://localhost:8080/actividades
GET https://localhost:8080/actividades/:id
GET https://localhost:8080/actividades/buscar?{id,titulo,horario,categoria}

	body: -

GET https://localhost:8080/usuarios?id
GET https://localhost:8080/usuarios/actividades
POST https://localhost:8080/login

	body: {
		"username": "string",
		"password": "string"
	}

POST https://localhost:8080/signup

	body: {
		"nombre": "string",
		"apellido": "string",
		"email": "string",
		"telefono": "string"
		"username": "string",
		"password": "string"
	}

POST https://localhost:8080/actividades Crear una actividad (Admin)

	body: {
  		"titulo": "string",
  		"descripcion": "string",
  		"cupo": "int",
  		"dia": "string",
  		"horario_inicio": "timestamp",
  		"horario_final": "timestamp",
  		"instructor_id": "int",
  		"categoria": "string"
}

PUT https://localhost:8080/actividades/:id
	body: {
		"titulo": "string",
  		"descripcion": "string",
  		"cupo": "int",
	}

DELETE https://localhost:8080/actividades/:id // TODO: preguntar porque al eliminar actividad, la inscripcion queda, ¿como se saca?
	

POST https://localhost:8080/inscripcion/:actividad_id

	header: autorization:bearer TOKEN       // TODO: preguntar si esta bien
	body: {
		"estado_inscripcion": "string",
		"fecha_inscripcion": "time",
  		"usuario_id": "int"
	}



*/

func login(ctx *gin.Context) {
	var loginDto dto.LoginDTO
	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// TODO: usar el servicio de usuario
	usuario, err := usuario.GetUsuarioByUsername(loginDto.User)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar usuario"})
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(loginDto.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
		log.Info("Contraseña incorrecta para", usuario.Username, "@", usuario.Password)
		return
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ocurrio un error en el servidor"})
		log.Error("La variable de entorno JWT_SECRET esta vacia")
		return
	}

	claims := jwt.MapClaims{
		"iss": "proyecto2025-morini-heredia",
		"exp": time.Now().Add(30 * time.Minute).Unix(),
		"sub": usuario.Username,
		"rol": usuario.Rol,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando token"})
		log.Error("Error generando el token:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": tokenString,
		"token_type":   "bearer",
		"expires_in":   1800, // en segundos
	})
}

func signup(ctx *gin.Context) {

}

func MapURLs() {
	router.GET("/actividades", actividad.GetAllActividades)
	router.GET("/actividades/:id", actividad.GetActividadById)
	router.GET("/actividades/buscar", actividad.GetActividadesByParams)

	router.POST("/login", login)
	router.POST("/signup", signup)
}
