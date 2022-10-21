package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateProductSuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Nama:      "hello",
		Deskripsi: "bandung",
		Harga:     230000,
		Kuantitas: 2,
	}

	update, err := service.UpdateProduct("1", req)
	assert.Empty(t, err)
	assert.NotNil(t, update)
}
