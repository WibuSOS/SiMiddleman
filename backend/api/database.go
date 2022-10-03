package api

import (
	"fmt"
	"github.com/WibuSOS/sinarmas/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func SetupDb() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	dbUrl := os.Getenv("DATABASE_URL")

	if os.Getenv("ENVIRONMENT") == "PROD" {
		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	} else {
		config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_SUPER_USER"),
			os.Getenv("DB_SUPER_PASSWORD"),
			os.Getenv("DB_ROOT"),
		)
		dbRoot, errs := gorm.Open(postgres.Open(config), &gorm.Config{})

		if errs != nil {
			return nil, fmt.Errorf("failed to connect database: %w", errs)
		}

		db = dbRoot.Exec(fmt.Sprintf("CREATE DATABASE %s;", os.Getenv("DB_NAME")))

		if db.Error != nil {
			fmt.Println("Unable to create DB, attempting to connect assuming it exists...")
			config = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_NAME"),
			)
			db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

			if err != nil {
				return nil, fmt.Errorf("failed to connect database: %w", err)
			}
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := db.AutoMigrate(&models.Todos{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, err
}
