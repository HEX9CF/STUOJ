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
func InitDatabase() {
	data := conf.Conf.DateBase
	var err error
	db, err = sql.Open("mysql", data.User+":"+data.Pwd+"@tcp("+data.Host+":"+data.Port+")/"+data.Name)
	log.Println("Connecting to MySQL: ", data.User+":"+data.Pwd+"@tcp("+data.Host+":"+data.Port+")/"+data.Name)

	if err != nil {
		log.Println("Open db error:", err)
		return
	} else {
		db.SetMaxIdleConns(data.MaxIdle)
		db.SetMaxOpenConns(data.MaxConn)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Error pinging the database:", err)
		return
	}

	log.Println("Successfully connected to MySQL!")
}
