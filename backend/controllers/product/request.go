package product

import "github.com/WibuSOS/sinarmas/utils/errors"

type DataRequest struct {
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
	Harga     uint   `json:"harga"`
	Kuantitas uint   `json:"kuantitas"`
}

func (r *DataRequest) ValidateReq() *errors.RestError {
	if r.Nama == "" {
		return errors.NewBadRequestError("Nama tidak boleh kosong")
	}

	if r.Harga <= 0 {
		return errors.NewBadRequestError("Harga tidak boleh kosong")
	}

	if r.Kuantitas <= 0 {
		return errors.NewBadRequestError("Kuantitas tidak boleh kosong")
	}

	if r.Deskripsi == "" {
		return errors.NewBadRequestError("Deskripsi tidak boleh kosong")
	}

	return nil
}
