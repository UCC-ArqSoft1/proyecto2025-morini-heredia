package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
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
	GetClaimsFromToken(tokenString string) (jwt.MapClaims, error)
}

var (
	IncorrectCredentialsError = errors.New("Credenciales incorrectas")

	UsuarioService IUsuarioService
	jwtSecret      string
)

func init() {
	UsuarioService = &usuarioService{}

	jwtSecret = os.Getenv("JWT_SECRET")
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
		"iss":        "proyecto2025-morini-heredia",
		"exp":        time.Now().Add(30 * time.Minute).Unix(),
		"username":   userdata.Username,
		"id_usuario": userdata.Id,
		"is_admin":   userdata.IsAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}

func (us *usuarioService) GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Error al obtener los claims")
	}

	return claims, nil
}
