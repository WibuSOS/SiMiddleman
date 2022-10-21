package authorization

import (
	"strconv"

	"github.com/WibuSOS/sinarmas/backend/helpers"
	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/WibuSOS/sinarmas/backend/utils/errors"

	"gorm.io/gorm"
)

type ServiceAuthorization interface {
	RoleAuthorize(role string) *errors.RestError
	RoomAuthorize(user_id float64, room_id string) *errors.RestError
}

type serviceAuthorization struct {
	db           *gorm.DB
	allowedRoles []string
}

func NewServiceAuthorization(db *gorm.DB, roles []string) *serviceAuthorization {
	return &serviceAuthorization{db: db, allowedRoles: roles}
}

func (s *serviceAuthorization) RoleAuthorize(role string) *errors.RestError {
	if !helpers.Contains(s.allowedRoles, role) {
		return errors.NewUnauthorized("unauthorized")
	}

	return nil
}

func (s *serviceAuthorization) RoomAuthorize(user_id float64, room_id string) *errors.RestError {
	var room models.Rooms

	roomID, err := strconv.ParseUint(room_id, 10, 64)
	if err != nil {
		return errors.NewBadRequestError("invalidRoomID")
	}

	res := s.db.Where("id = ?", roomID).First(&room)
	if res.Error != nil {
		return errors.NewBadRequestError("roomNotFound")
	}

	isPenjual := uint(user_id) == room.PenjualID
	var isPembeli bool
	if room.PembeliID == nil {
		isPembeli = false
	} else {
		isPembeli = uint(user_id) == *room.PembeliID
	}

	if !(isPenjual || isPembeli) {
		return errors.NewUnauthorized("unauthorized")
	}

	return nil
}
