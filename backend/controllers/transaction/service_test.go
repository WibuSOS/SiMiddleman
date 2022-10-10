package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPaymentDetailsServiceSuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	idRoom := newTestCreateRoomWithProduct(t)

	res, err := service.GetPaymentDetails(int(idRoom))
	assert.Nil(t, err)
	assert.Greater(t, res.Total, 0)
}

func TestGetPaymentDetailsServiceEmptyProduct(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	idRoom := newTestCreateRoomWithProduct(t)

	res, err := service.GetPaymentDetails(int(idRoom))
	assert.Nil(t, err)
	assert.Equal(t, 0, res.Total)
}

func TestGetPaymentDetailsServiceRoomNotFound(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	idRoom := newTestCreateRoomWithProduct(t)

	_, err := service.GetPaymentDetails(int(idRoom))
	assert.NotNil(t, err)
	assert.Equal(t, "Room Not Found", err.Message)
	assert.Equal(t, 400, err.Status)
	assert.Equal(t, "Bad_Request", err.Error)
}
