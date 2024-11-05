package db

import (
	"STUOJ/conf"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	SqlDb *sql.DB
	Db    *gorm.DB
)

// 初始化数据库
func InitDatabase() error {
	var err error
	config := conf.Conf.Datebase

	dsn := config.User + ":" + config.Pwd + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println("Connecting to MySQL:", dsn)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect database!")
		return err
	}

	SqlDb, err = Db.DB()
	if err != nil {
		log.Println("Failed to get sql.SqlDb!")
		return err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	SqlDb.SetMaxIdleConns(config.MaxIdle)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	SqlDb.SetMaxOpenConns(config.MaxConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	SqlDb.SetConnMaxLifetime(time.Hour)

	/*		err = SqlDb.Ping()
			if err != nil {
				log.Println("Error pinging the database!")
				return err
			}
	*/
	log.Println("Database init success!")
	return nil
}
