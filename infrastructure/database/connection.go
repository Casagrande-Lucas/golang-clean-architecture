package database

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

type Database interface {
	Connect() (*gorm.DB, error)
}

type MySQLDatabase struct {
	DSN string
}

func (m *MySQLDatabase) Connect() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(m.DSN), &gorm.Config{})
}

type PostgresDatabase struct {
	DSN string
}

func (p *PostgresDatabase) Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(p.DSN), &gorm.Config{})
}

func GetDBInstance() *gorm.DB {
	once.Do(func() {
		var dbConn Database
		dbType := os.Getenv("DB_TYPE")
		dsn := os.Getenv("DB_DSN")

		switch dbType {
		case "mysql":
			dbConn = &MySQLDatabase{DSN: dsn}
		case "postgres":
			dbConn = &PostgresDatabase{DSN: dsn}
		default:
			log.Fatal("Unsupported database type:", dbType)
		}

		db, err := dbConn.Connect()
		if err != nil {
			log.Fatal("failed to connect database:", err)
		}
		dbInstance = db
	})
	return dbInstance
}
