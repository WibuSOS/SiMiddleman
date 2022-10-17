package api

import (
	"fmt"
	"log"
	"os"

	"github.com/WibuSOS/sinarmas/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func errorDbConn(err error) error {
	return fmt.Errorf("failed to connect database: %w", err)
}

func callDbDev() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_SUPER_USER"), os.Getenv("DB_SUPER_PASSWORD"), os.Getenv("DB_ROOT"))
	dbRoot, errRoot := gorm.Open(postgres.Open(config), &gorm.Config{})

	if errRoot != nil {
		return nil, errorDbConn(errRoot)
	}

	db = dbRoot.Exec(fmt.Sprintf("CREATE DATABASE %s ;", os.Getenv("DB_NAME")))

	if db.Error != nil {
		log.Println("Unable to create DB, attempting to connect assuming it exists...")
		config = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
		db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	}

	if err != nil {
		return nil, errorDbConn(err)
	}

	return db, nil
}

func callDb() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	env := os.Getenv("ENVIRONMENT")

	if env == "PROD" || env == "STAGING" {
		dbUrl := os.Getenv("DATABASE_URL")
		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	}

	if err != nil {
		return nil, errorDbConn(err)
	}

	if db != nil {
		return db, nil
	}

	db, err = callDbDev()

	if err != nil {
		return nil, errorDbConn(err)
	}

	return db, nil
}

func checkDbConn(db *gorm.DB) (*gorm.DB, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errorDbConn(err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, errorDbConn(err)
	}

	return db, nil
}

func SetupDb() (*gorm.DB, error) {
	db, err := callDb()

	if err != nil {
		return nil, err
	}

	db, err = checkDbConn(db)
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Users{}, &models.Rooms{}, &models.Products{}, &models.Transactions{}); err != nil {
		return nil, errorDbConn(err)
	}

	return db, nil
}
