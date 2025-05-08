package model

import "time"

type Usuario struct {
	Id            int
	Nombre        string
	Apellido      string
	Email         string
	UserName      string
	Password      string
	Rol           string
	FechaRegistro time.Time

	Actividades Actividades
}

type Usuarios []Usuario
