package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"proyecto-integrador/clients/usuario"
	"proyecto-integrador/model"
	"time"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

type usuarioService struct{}

type IUsuarioService interface {
	GenerateToken(username string, password string) (string, error)
}

var (
	IncorrectCredentialsError = errors.New("Credenciales incorrectas")

	UsuarioService IUsuarioService
)

func init() {
	UsuarioService = &usuarioService{}
}

func (us *usuarioService) GenerateToken(username string, password string) (string, error) {
	var userdata model.Usuario = usuario.GetUsuarioByUsername(username)

	hashedPassword := sha256.Sum256([]byte(password))
	if hex.EncodeToString(hashedPassword[:]) != userdata.Password {
		log.Debugf("Contraseña incorrecta para el usuario %s@%s\n", username, password)
		log.Debugf("Hash ingresado: %s", hex.EncodeToString(hashedPassword[:]))
		return "", IncorrectCredentialsError
	}

	claims := jwt.MapClaims{
		"iss": "proyecto2025-morini-heredia",
		"exp": time.Now().Add(30 * time.Minute).Unix(),
		"sub": userdata.Username,
		"rol": userdata.IsAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Error("La variable de entorno JWT_SECRET esta vacia")
		return "", fmt.Errorf("La variable de entorno JWT_SECRET esta vacia")
	}

	return token.SignedString([]byte(secret))
}
