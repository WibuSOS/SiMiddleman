package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func TestDb(t *testing.T) {
// 	var db *gorm.DB
// 	var err error
// 	DB_USER := "postgres"
// 	DB_PASSWORD := "ikangurami"
// 	DB_NAME := "simiddleman"
// 	DB_ROOT := "postgres"
// 	DB_HOST := "localhost"
// 	DB_PORT := "5432"
// 	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		DB_HOST,
// 		DB_PORT,
// 		DB_USER,
// 		DB_PASSWORD,
// 		DB_ROOT,
// 	)
// 	dbRoot, errs := gorm.Open(postgres.Open(config), &gorm.Config{})
// 	assert.NoError(t, errs)
// 	assert.NotNil(t, dbRoot)

// 	db = dbRoot.Exec(fmt.Sprintf("SELECT 'CREATE DATABASE %s' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '%s')", DB_NAME, DB_NAME))
// 	assert.NoError(t, db.Error)
// 	assert.NotNil(t, db)

// 	sqlDB, err := db.DB()
// 	assert.NoError(t, err)
// 	assert.NotNil(t, sqlDB)

// 	err = sqlDB.Ping()
// 	assert.NoError(t, err)

// 	err = db.AutoMigrate(models.Users{})
// 	assert.NoError(t, err)
// }

func TestServer(t *testing.T) {
	config := "host=localhost user=postgres password=ikangurami dbname=simiddleman port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	server := MakeServer(db)
	// server.RunServer()
	assert.NotNil(t, server)

	// os.Setenv("asd", "8080")
	// server.RunServer()
	// assert.Error(t, server.DB.Error)
}