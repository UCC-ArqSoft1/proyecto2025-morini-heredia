package model

import "time"

type Actividad struct {
	Id          uint   `gorm:"primaryKey;autoIncrement"`
	Titulo      string `gorm:"type:varchar(100);not null"`
	Descripcion string `gorm:"type:text"`
	Cupo        uint   `gorm:"type:int;not null"`
	// TODO: esto conviene que sea multivaluado, porque sino tendríamos que crear varias actividades para una sola actividad
	Dia            string    `gorm:"type:enum('Lunes','Martes','Miercoles','Jueves','Viernes','Sabado','Domingo');not null'"`
	Horario_inicio time.Time `gorm:"type:timestamp"`
	Horario_final  time.Time `gorm:"type:timestamp"`
	Instructor     string    `gorm:"type:varchar(100);not null"`

	Categoria string `gorm:"type:varchar(24)"`
}

type Actividades []Actividad
