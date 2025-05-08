package db

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"sync"
)

var (
	db   *gorm.DB
	once sync.Once

	db_user   = os.Getenv("DB_USER")
	db_pass   = os.Getenv("DB_PASS")
	db_host   = os.Getenv("DB_HOST")
	db_schema = os.Getenv("DB_SCHEMA")
)

func GetInstance() *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_host, db_schema)

		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Error al conectar a la base de datos: %v", err)
		}
		log.Info("Conexión a base de datos establecida")
	})

	return db
}

func StartDbEngine() {
	// creamos la conexión
	GetInstance()
}
