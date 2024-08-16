package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func NewDatabase(config *viper.Viper) *gorm.DB {
	var gormDB *gorm.DB

	connection := config.GetString("DB_CONNECTION")
	username := config.GetString("DB_USERNAME")
	password := config.GetString("DB_PASSWORD")
	host := config.GetString("DB_HOST")
	port := config.GetInt("DB_PORT")
	database := config.GetString("DB_DATABASE")
	idleConnection := config.GetInt("DB_IDLE_CONNECTION")
	maxConnection := config.GetInt("DB_MAX_CONNECTION")
	maxLifeTimeConnection := config.GetInt("DB_MAX_LIFE_TIME_CONNECTION")

	// mysql
	if connection == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username,
			password,
			host,
			port,
			database,
		)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		conn, err := db.DB()
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		conn.SetMaxIdleConns(idleConnection)
		conn.SetMaxOpenConns(maxConnection)
		conn.SetConnMaxLifetime(time.Duration(maxLifeTimeConnection) * time.Second)

		gormDB = db
	}

	return gormDB
}
