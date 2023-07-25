package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	Driver   string `json:"dbDriver"`
	Host     string `json:"dbHost"`
	Port     string `json:"dbPort"`
	Name     string `json:"dbName"`
	User     string `json:"dbUser"`
	Password string `json:"dbPassword"`
}

func NewDatabaseConnection() (*gorm.DB, error) {

	var config = DatabaseConfig{
		Driver:   "postgres",
		Host:     "127.0.0.1",
		Port:     "5432",
		Name:     "ecommerce",
		User:     "postgres",
		Password: "123456",
	}

	var dsn string
	var db *gorm.DB

	switch config.Driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Name)
		db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Paris", config.Host, config.User, config.Password, config.Name, config.Port)
		db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", config.Driver)
	}

	return db, nil
}
