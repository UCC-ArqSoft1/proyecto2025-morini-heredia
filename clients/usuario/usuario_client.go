package usuario

import (
	"proyecto-integrador/db"

	"gorm.io/gorm"
)

var db_conn *gorm.DB = db.GetInstance()
