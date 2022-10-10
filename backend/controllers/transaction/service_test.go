package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceUpdateStatus(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	err := service.UpdateStatusDelivery("1")
	assert.Nil(t, err)

}

func TestServiceUpdateStatusError(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	err := service.UpdateStatusDelivery("asd")
	assert.NotNil(t, err)
	assert.Equal(t, "WHERE conditions required", err.Message)
	assert.Equal(t, 400, err.Status)

}
