package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// CustomTime es un tipo personalizado que siempre usa el año 2024 al guardar en la base de datos
type CustomTime time.Time

func (ct CustomTime) Format(layout string) string {
	return time.Time(ct).Format(layout)
}

func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		*ct = CustomTime(time.Time{})
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		// Asegurarse de que siempre usamos el año 2024
		t := time.Date(2024, v.Month(), v.Day(), v.Hour(), v.Minute(), v.Second(), v.Nanosecond(), v.Location())
		*ct = CustomTime(t)
		return nil
	default:
		return fmt.Errorf("cannot scan %T into CustomTime", value)
	}
}

func (ct CustomTime) Value() (driver.Value, error) {
	t := time.Time(ct)
	// Asegurarse de que siempre usamos el año 2024
	return time.Date(2024, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location()), nil
}

// String implementa la interfaz Stringer para mostrar la hora en formato legible
func (ct CustomTime) String() string {
	return time.Time(ct).Format("15:04")
}

type Actividad struct {
	Id            uint       `gorm:"column:id_actividad;primaryKey;autoIncrement"`
	Titulo        string     `gorm:"type:varchar(50);not null"`
	Descripcion   string     `gorm:"type:varchar(255)"`
	Cupo          uint       `gorm:"type:int;not null"`
	Dia           string     `gorm:"type:enum('Lunes','Martes','Miercoles','Jueves','Viernes','Sabado','Domingo');not null"`
	HorarioInicio CustomTime `gorm:"column:horario_inicio;type:time;not null"`
	HorarioFinal  CustomTime `gorm:"column:horario_final;type:time;not null"`
	FotoUrl       string     `gorm:"column:foto_url;type:varchar(511);not null"`
	Instructor    string     `gorm:"type:varchar(50);not null"`
	Categoria     string     `gorm:"type:varchar(40);not null"`

	Inscripciones Inscripciones `gorm:"foreignKey:IdActividad"`
}

type Actividades []Actividad

type ActividadVista struct {
	Id            uint       `gorm:"column:id_actividad;primaryKey;autoIncrement"`
	Titulo        string     `gorm:"type:varchar(50);not null"`
	Descripcion   string     `gorm:"type:varchar(255)"`
	Cupo          uint       `gorm:"type:int;not null"`
	Dia           string     `gorm:"type:enum('Lunes','Martes','Miercoles','Jueves','Viernes','Sabado','Domingo');not null"`
	HorarioInicio CustomTime `gorm:"column:horario_inicio;type:time;not null"`
	HorarioFinal  CustomTime `gorm:"column:horario_final;type:time;not null"`
	FotoUrl       string     `gorm:"column:foto_url;type:varchar(511);not null"`
	Instructor    string     `gorm:"type:varchar(50);not null"`
	Categoria     string     `gorm:"type:varchar(40);not null"`
	Lugares       uint       `gorm:"column:lugares"` // Campo calculado de la vista
}

type ActividadesVista []ActividadVista

func (ActividadVista) TableName() string {
	return "actividads_lugares"
}
