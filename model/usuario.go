package model

import "time"

type Usuario struct {
	Id            int
	Nombre        string
	Apellido      string
	Email         string
	UserName      string
	Password      string `gorm:"not null" json:"-"` // no expone la contraseña en el JSON
	Rol           string
	FechaRegistro time.Time

	Actividades Actividades
}

type Usuarios []Usuario
