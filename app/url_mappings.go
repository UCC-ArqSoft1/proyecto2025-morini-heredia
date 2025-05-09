package app

import (
	"net/http"
	"os"
	"proyecto-integrador/clients/actividad"
	"proyecto-integrador/clients/usuario"
	"proyecto-integrador/dto"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
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
	https://localhost:8080/usuarios?id
	https://localhost:8080/usuarios/actividades
	https://localhost:8080/login
	https://localhost:8080/signup
*/

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
	var dto dto.LoginDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	usuario, err := usuario.GetUsuarioByUsername(dto.User)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar usuario"})
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(dto.Password)); err != nil {
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
	// TODO: declarar endpoints de la app
	router.GET("/actividades", getActividades)
	router.GET("/actividades/:id", getActividadById)
	router.GET("/usuario/actividades", getActividadesUsuario)

	router.POST("/login", login)
	router.POST("/signup", signup)
}
