package model

type Categoria struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Nombre string `gorm:"type:varchar(100);not null"`
}

type Categorias []Categoria
