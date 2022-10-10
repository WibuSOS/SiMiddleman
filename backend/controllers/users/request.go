package users

import (
	"regexp"

	"github.com/WibuSOS/sinarmas/utils/errors"
)

type DataRequest struct {
	Nama     string `json:"nama" binding:"required"`
	NoHp     string `json:"noHp" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	NoRek    string `json:"noRek" binding:"required"`
}

func (req *DataRequest) ValidateReq() *errors.RestError {
	if req.Nama == "" || len(req.Nama) > 30 {
		return errors.NewBadRequestError("nama tidak memenuhi syarat")
	}

	regex := regexp.MustCompile(`^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,7}$`)
	if matched := regex.MatchString(req.NoHp); req.NoHp == "" || !matched || len(req.NoHp) > 18 {
		return errors.NewBadRequestError("no hp tidak memenuhi syarat")
	}

	regex = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if matched := regex.MatchString(req.Email); req.Email == "" || !matched || len(req.Email) > 30 {
		return errors.NewBadRequestError("email tidak memenuhi syarat")
	}

	if req.Password == "" || len(req.Password) < 8 || len(req.Password) > 18 {
		return errors.NewBadRequestError("password tidak memenuhi syarat")
	}

	regex = regexp.MustCompile(`^[0-9]{4,18}$`)
	if matched := regex.MatchString(req.NoRek); req.NoRek == "" || !matched || len(req.NoRek) < 4 || len(req.NoRek) > 18 {
		return errors.NewBadRequestError("no rek tidak memenuhi syarat")
	}

	return nil
}
