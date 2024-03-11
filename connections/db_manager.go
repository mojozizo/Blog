package connections

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbManager() error {
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
