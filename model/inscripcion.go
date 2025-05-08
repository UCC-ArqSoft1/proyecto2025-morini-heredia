package model

type Inscripcion struct {
	Id        int
	Usuario   Usuario
	Actividad Actividad
}

type Inscripciones []Inscripcion
