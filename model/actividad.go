package model

import "time"

type Actividad struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	Titulo        string    `gorm:"type:varchar(100);not null"`
	Descripcion   string    `gorm:"type:text"`
	Cupo          uint      `gorm:"type:int;not null"`
	Dia           string    `gorm:"type:enum('Lunes','Martes','Miercoles','Jueves','Viernes','Sabado','Domingo');not null'"`
	Categoria     string    `gorm:"type:varchar(24);not null"`
	HorarioInicio time.Time `gorm:"type:timestamp"`
	HorarioFinal  time.Time `gorm:"type:timestamp"`

	InstructorID uint       `gorm:"not null"` // Clave foránea para Instructor
	Instructor   Instructor `gorm:"foreignKey:InstructorId"`
	// Usuarios se maneja a través de Inscripcion (muchos-a-muchos)
}

type Actividades []Actividad
