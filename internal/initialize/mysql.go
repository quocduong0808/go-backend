package initialize

import (
	"fmt"
	"go/go-backend-api/global"
	"go/go-backend-api/internal/po"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
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
	if config.AutoCreate {
		//migrateTable()
		genModels()
		global.Logger.Info("auto migrate tables/models successfully")
	}
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

func migrateTable() {
	err := global.MyDB.AutoMigrate(&po.User{}, &po.Role{})
	global.HandleErrorPanic(err, "error when auto migrate table")
}

func genModels() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",                                                 // output path
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.MyDB) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.GenerateAllTable()

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}
