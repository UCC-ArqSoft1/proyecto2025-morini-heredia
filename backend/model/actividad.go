package model

import "time"

type Actividad struct {
	Id            uint      `gorm:"column:id_actividad;primaryKey;autoIncrement"`
	Titulo        string    `gorm:"type:varchar(50);not null"`
	Descripcion   string    `gorm:"type:varchar(255)"`
	Cupo          uint      `gorm:"type:int;not null"`
	Dia           string    `gorm:"type:enum('Lunes','Martes','Miercoles','Jueves','Viernes','Sabado','Domingo');not null'"`
	HorarioInicio time.Time `gorm:"column:horario_inicio;type:time;not null"`
	HorarioFinal  time.Time `gorm:"column:horario_final;type:time;not null"`
	FotoUrl       string    `gorm:"column:foto_url;type:varchar(511);not null"`
	Instructor    string    `gorm:"type:varchar(50);not null"`
	Categoria     string    `gorm:"type:varchar(40);not null"`

	Inscripciones Inscripciones `gorm:"foreignKey:IdActividad"`
}

type Actividades []Actividad
