package db

import (
	"database/sql"
	"gorm.io/gorm"
)

var (
	Db    *gorm.DB
	SqlDb *sql.DB
)
