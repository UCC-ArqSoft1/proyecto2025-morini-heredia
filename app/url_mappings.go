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
INICIO DE SESION/CREACION DE USUARIO:
	POST https://localhost:8080/login
		http_status: 201, 400, 401, 500
		body: {
			"username": "string",
			"password": "string"
		}
		respose: {
			"access_token": "TOKEN",
		}

	POST https://localhost:8080/signup
		http_status: 201, 400, 500
		body: {
			"nombre": "string",
			"apellido": "string",
			"email": "string",
			"telefono": "string"
			"username": "string",
			"password": "string"
		}


ADMINISTRACION DE LOS MODELOS:

	GET https://localhost:8080/actividades
		http_status: 200, 500
		response: {
			"actividades": [
				{
					"id": "int",
					"titulo": "string", ...
				}
			]
		}
	GET https://localhost:8080/actividades/buscar?{id,titulo,horario,categoria}
		http_status: 200, 500
		response: {
			"actividades": [
				{
					"id": "int",
					"titulo": "string", ...
				}
			]
		}
	GET https://localhost:8080/actividades/:id_actividad
		http_status: 200, 404, 400, 500
		response:
			{
				"id": "int",
				"titulo": "string", ...
			}
	// TODO: según la consigna para consultar actividades de un socio no se requiere autenticación, preguntar si está bien (sería muy raro)
	GET https://localhost:8080/usuarios/actividades
		http_status: 200, 401, 500
		header: autorization:bearer TOKEN
		response: {
			"actividades": [
				{
					"id": "int",
					"titulo": "string", ...
				}
			]
		}
	POST https://localhost:8080/inscripciones/:actividad_id
		status: 201, 404, 401, 500
		header: autorization:bearer TOKEN       //TODO: preguntar si esta bien (buscar al usuario usando contenido del token) y consultar como buscar al usuario
		body: {
			"estado_inscripcion": "string"
		}
	DELETE https://localhost:8080/inscripciones/:actividad_id  // Permitir que un usuario elimine su inscripción en una actividad
		http_status: 204, 404, 401, 500
		header: autorization:bearer TOKEN
		body: {
			"usuario_id": "int"
		}


ENDPOINTS PARA EL ADMINISTRADOR:

	Crear una actividad
	POST https://localhost:8080/actividades (Admin)
		http_status: 201, 400, 401, 403, 500
		header: autorization:bearer TOKEN
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

	Actualizar una actividad
	PUT https://localhost:8080/actividades/:id (Admin)
		http_status: 200, 400, 401, 403, 500
		header: autorization:bearer TOKEN
		body: {
			"titulo": "string",
			"descripcion": "string",
			"cupo": "int",
		}
		response: <BODY>

	Borrar una actividad
	DELETE https://localhost:8080/actividades/:id (Admin) //TODO: preguntar porque al eliminar actividad, la inscripcion queda, ¿como se saca?
		http_status: 204, 404, 401, 403, 500 // TODO: no se si iba 404 con el metodo DELETE
		header: autorization:bearer TOKEN
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
