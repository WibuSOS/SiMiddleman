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
		return errors.NewBadRequestError("invalidEmail")
	}

	if r.Password == "" {
		return errors.NewBadRequestError("invalidPassword")
	}

	return nil
}
