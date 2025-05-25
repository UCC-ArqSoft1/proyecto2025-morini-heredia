package app

import (
	"net/http"
	"proyecto-integrador/services"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func JWTValidation(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "debes especificar el header 'Authorization' con tu token"})
		return
	}

	parts := strings.SplitN(auth, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no autorizado"})
		return
	}

	uid, err := services.UsuarioService.ValidateToken(parts[1])
	if err != nil {
		log.Debug(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no autorizado"})
		return
	}

	ctx.Set("id_usuario", uid)
	ctx.Next()
}

func MapMiddewares() {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}
