package model

type Actividad struct {
	Id          int
	Titulo      string
	Descripcion string
	Cupo        int
	// instructor
	// categoria

	Usuarios Usuarios
}

type Actividades []Actividad
