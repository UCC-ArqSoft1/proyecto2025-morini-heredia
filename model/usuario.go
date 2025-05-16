package model

import "time"

type Usuario struct {
	Id             uint      `gorm:"primaryKey;autoIncrement"`
	Nombre         string    `gorm:"type:varchar(100);not null"`
	Apellido       string    `gorm:"type:varchar(100);not null"`
	Username       string    `gorm:"type:varchar(50);unique;not null"`
	Password       string    `gorm:"type:varchar(255);not null"` // no expone la contraseña en el JSON
	Rol            string    `gorm:"type:enum('socio','admin');not null;default:'socio'"`
	Fecha_registro time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`

	// Relación muchos-a-muchos (a través de Inscripcion)
	Actividades []Actividad `gorm:"many2many:inscripciones;"`
}

type Usuarios []Usuario
