package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSetupDb(t *testing.T) {
	var db *gorm.DB
	var err error

	db, err = SetupDb()
	assert.NotNil(t, db)
	assert.NoError(t, err)
}
