package product

import "github.com/WibuSOS/sinarmas/utils/errors"

type DataRequest struct {
	Nama      string `json:"nama"`
	Harga     int    `json:"harga"`
	Kuantitas int    `json:"kuantitas"`
	Deskripsi string `json:"deskripsi"`
}

func (r *DataRequest) ValidateReq() *errors.RestError {
	if r.Nama == "" {
		return errors.NewBadRequestError("Nama tidak boleh kosong")
	}

	if r.Deskripsi == "" {
		return errors.NewBadRequestError("Deskripsi tidak boleh kosong")
	}
	return nil
}
