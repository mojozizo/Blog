package connections

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SQLDbManager() error {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		return fmt.Errorf("DB_URL environment variable is not set")
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	return nil
}

var PgDB *gorm.DB

func PostgreSQLDbManager() error {
	dsn := os.Getenv("POSTGRE_DB_URL")
	if dsn == "" {
		return fmt.Errorf("POSTGRE_DB_URL environment variable is not set")
	}

	var err error
	PgDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	return nil
}
