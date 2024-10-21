package db

import (
	"STUOJ/conf"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	db *sql.DB
)

// InitDatabase 函数用于初始化数据库连接
func InitDatabase() error {
	data := conf.Conf.DateBase
	var err error

	dsn := data.User + ":" + data.Pwd + "@tcp(" + data.Host + ":" + data.Port + ")/" + data.Name
	db, err = sql.Open("mysql", dsn)
	log.Println("Connecting to MySQL:", dsn)

	if err != nil {
		log.Println("Open database error!")
		return err
	}
	db.SetMaxIdleConns(data.MaxIdle)
	db.SetMaxOpenConns(data.MaxConn)
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Error pinging the database!")
		return err
	}

	log.Println("Successfully connected to MySQL!")
	return nil
}
