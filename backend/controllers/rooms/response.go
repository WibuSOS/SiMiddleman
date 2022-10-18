package rooms

import (
	"time"

	"github.com/WibuSOS/sinarmas/backend/controllers/users"
	"github.com/WibuSOS/sinarmas/backend/models"
)

type DataResponse struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	PenjualID uint                `json:"penjualID"`
	PembeliID *uint               `json:"pembeliID"`
	RoomCode  string              `json:"roomCode"`
	Status    string              `json:"status"`
	Product   *models.Products    `json:"product"`
	Penjual   *users.DataResponse `json:"penjual"`
	Pembeli   *users.DataResponse `json:"pembeli"`
}
