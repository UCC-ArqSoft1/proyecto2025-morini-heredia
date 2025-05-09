package model

import "time"

type Actividad struct {
	Id             uint      `gorm:"primaryKey;autoIncrement"`
	Titulo         string    `gorm:"type:varchar(100);not null"`
	Descripcion    string    `gorm:"type:text"`
	Cupo           uint      `gorm:"type:int;not null"`
	Dia            string    `gorm:"type:enum('Lunes','Martes','Miercoles','Jueves','Viernes','Sabado','Domingo');not null'"`
	Horario_inicio time.Time `gorm:"type:timestamp"`
	Horario_final  time.Time `gorm:"type:timestamp"`

	Instructor_id uint       `gorm:"not null"` // Clave foránea para Instructor
	Instructor    Instructor `gorm:"foreignKey:instructor_id"`
	// Usuarios se maneja a través de Inscripcion (muchos-a-muchos)c

	Categoria_id uint      `gorm:"not null"` // Clave foránea para Instructor
	Categoria    Categoria `gorm:"foreignKey:categoria_id"`
}

type Actividades []Actividad
