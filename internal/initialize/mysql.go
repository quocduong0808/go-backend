package initialize

import (
	"fmt"
	"go/go-backend-api/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	config := global.Config.Mysql
	uri := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(uri, config.Username, config.Password, config.Host, config.Port, config.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	global.HandleErrorPanic(err, "error when init connection to database")
	initMysqlPool(db)
	global.MyDB = db
	global.Logger.Info("create database connection successfully")
}

func initMysqlPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	config := global.Config.Mysql
	global.HandleErrorPanic(err, "error when init database connection pool")
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(config.MaxIdleCons)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(config.MaxOpenCons)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifeTime) * time.Second)
}
