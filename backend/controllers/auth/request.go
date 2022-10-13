package auth

import (
	"github.com/WibuSOS/sinarmas/backend/utils/errors"
)

type DataRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *DataRequest) Validation() *errors.RestError {
	if r.Email == "" {
		return errors.NewBadRequestError("Invalid email")
	}

	if r.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}

	return nil
}
