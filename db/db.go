package db

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"STUOJ/conf"
)

var(
	db *sql.DB
)
// InitDatabase 函数用于初始化数据库连接
func InitDatabase() {
	data:=conf.Conf.DateBase
	db, err := sql.Open("mysql", data.User+":"+data.Pwd+"@tcp("+data.Host+":"+data.Port+")/"+data.Name)
	if err != nil {
		fmt.Println("open db error:", err)
		return
	} else {
		db.SetMaxIdleConns(data.MaxIdle)
		db.SetMaxOpenConns(data.MaxConn)
	}
	defer db.Close()

	err = db.Ping()
    if err != nil {
        fmt.Println("Error pinging the database:", err)
        return
    }

    fmt.Println("Successfully connected to MySQL!")
}
