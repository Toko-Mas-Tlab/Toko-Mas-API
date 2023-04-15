package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type credentialDB struct {
	dbHost     string
	dbUser     string
	dbName     string
	dbPassword string
	dbPort     string
}

func ConnectionDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		// return "godotenv failed", err
		return nil, err
	}

	cred := credentialDB{
		dbHost:     os.Getenv("POSTGRES_HOST"),
		dbUser:     os.Getenv("POSTGRES_USER"),
		dbName:     os.Getenv("POSTGRES_DB"),
		dbPassword: os.Getenv("POSTGRES_PASSWORD"),
		dbPort:     os.Getenv("POSTGRES_PORT"),
	}

	dsn := "host=" + cred.dbHost + " user=" + cred.dbUser + " password=" + cred.dbPassword + " dbname=" + cred.dbName + " port=" + cred.dbPort + " sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// return "Failed Connect to Database", err
		return nil, err

	}

	return DB, nil
}
