package database_connection

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ahmadirfaan/match-nearby-app-rest/app"
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDb() *gorm.DB {
	app := app.Init()
	maxIdleConn := app.Config.DBMaxIdleConnections
	maxConn := app.Config.DBMaxConnections
	maxLifetimeConn := app.Config.DBMaxLifetimeConnections
	db_user := app.Config.DBUsername
	db_pass := app.Config.DBPassword
	db_host := app.Config.DBHost
	db_port := app.Config.DBPort
	db_database := app.Config.DBName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port, db_database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default,
		SkipDefaultTransaction: true,
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: false,
			NameReplacer:  strings.NewReplacer("ID", "id"),
			NoLowerCase:   false,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	InitCreateTable(db)
	log.Println("database connect success")
	return db

}

func InitCreateTable(db *gorm.DB) {

	db.Debug().AutoMigrate(
		&database.Users{},
		&database.Swipes{},
		&database.Profiles{},
		&database.Subscriptions{},
	)

}
