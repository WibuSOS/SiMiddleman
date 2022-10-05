package product

import (
	"github.com/WibuSOS/sinarmas/utils/errors"
)

type DataRequest struct {
	Nama      string `json:"nama"`
	Harga     int    `json:"harga"`
	Kuantitas int    `json:"kuantitas"`
	Deskripsi string `json:"deskripsi"`
}

func (r *DataRequest) Validation() *errors.RestError {
	if r.Nama == "" {
		return errors.NewBadRequestError("Nama harus diisi")
	}

	if r.Harga <= 0 {
		return errors.NewBadRequestError("Salah input Harga")
	}

	if r.Kuantitas <= 0 {
		return errors.NewBadRequestError("Salah input kunatitas")
	}

	if r.Deskripsi == "" {
		return errors.NewBadRequestError("Deskripsi harus diisi")
	}

	return nil
}
