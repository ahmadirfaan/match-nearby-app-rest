package databaseconnection

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

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
	databaseUsername := app.Config.DBUsername
	dbPassword := app.Config.DBPassword
	databaseHost := app.Config.DBHost
	databasePort := app.Config.DBPort
	databaseName := app.Config.DBName
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", databaseHost, databaseUsername, dbPassword, databaseName, databasePort)
	log.Info("dsn format : " + dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default,
		SkipDefaultTransaction: true,
		NamingStrategy: &schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: false,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		log.Info("error open postgres : " + dsn)
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
	InitEnums(db)
	db.Debug().AutoMigrate(
		&database.Users{},
		&database.Swipes{},
		&database.Profiles{},
		&database.Subscriptions{},
	)

}

func InitEnums(db *gorm.DB) {
	log.Info("run Init Enums")
	db.Exec(`
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
			CREATE TYPE gender_enum AS ENUM ('MALE', 'FEMALE');
		END IF;
	END$$;
`)
}
