package transaction

import (
	"testing"

	"github.com/WibuSOS/sinarmas/models"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(&models.Rooms{})
	assert.NoError(t, err)

	db.Create(&models.Rooms{
		PenjualID: 1,
	})

	return db
}

func TestUpdateStatusSuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	err := repo.UpdateStatusDelivery("1")
	assert.Empty(t, err)

}

func TestUpdateStatusErrorConvert(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	err := repo.UpdateStatusDelivery("ase")
	assert.NotNil(t, err)

}
