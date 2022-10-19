package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestConnectDB(t *testing.T) {
	var db *gorm.DB
	var err error
	os.Setenv("ENVIRONMENT", "ENV")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_SUPER_USER", "postgres")
	os.Setenv("DB_SUPER_PASSWORD", "ikangurami")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "ikangurami")
	os.Setenv("DB_NAME", "simiddleman")
	os.Setenv("DB_ROOT", "postgres")
	os.Setenv("DATABASE_URL", "database.db")
	db, err = SetupDb()
	assert.NoError(t, err)
	assert.NotNil(t, db)
	assert.Nil(t, db)
}

func TestConnectDBErr(t *testing.T) {
	var db *gorm.DB
	var err error
	os.Setenv("ENVIRONMENT", "ENV")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_SUPER_USER", "postgres123")
	os.Setenv("DB_SUPER_PASSWORD", "ikangurami")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "ikangurami")
	os.Setenv("DB_NAME", "simiddleman")
	os.Setenv("DB_ROOT", "postgres")
	os.Setenv("DATABASE_URL", "database.db")
	db, err = SetupDb()
	assert.Error(t, err)
	assert.Nil(t, db)
}

func TestConnectDBErr2(t *testing.T) {
	var db *gorm.DB
	var err error
	os.Setenv("ENVIRONMENT", "ENV")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_SUPER_USER", "postgres")
	os.Setenv("DB_SUPER_PASSWORD", "ikangurami")
	os.Setenv("DB_USER", "postgres123")
	os.Setenv("DB_PASSWORD", "ikangurami")
	os.Setenv("DB_NAME", "simiddleman")
	os.Setenv("DB_ROOT", "postgres")
	os.Setenv("DATABASE_URL", "database.db")
	db, err = SetupDb()
	assert.Error(t, err)
	assert.Nil(t, db)
}

func TestConnectDBProd(t *testing.T) {
	var db *gorm.DB
	var err error
	os.Setenv("DATABASE_URL", "database.db")
	os.Setenv("ENVIRONMENT", "PROD")
	db, err = SetupDb()
	assert.NoError(t, err)
	assert.NotNil(t, db)
}
