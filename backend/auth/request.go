package auth

import (
	"github.com/WibuSOS/SiMiddleman/utils/errors"
)

type DataRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *DataRequest) Validation() *errors.RestError {
	if r.Email == "" {
		return errors.NewBadRequestError("Invalid email")
	}

	if r.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
