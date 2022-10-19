package product

import "github.com/WibuSOS/sinarmas/backend/utils/errors"

type DataRequest struct {
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
	Harga     uint   `json:"harga"`
	Kuantitas uint   `json:"kuantitas"`
}

func (r *DataRequest) ValidateReq() *errors.RestError {
	if r.Nama == "" {
		return errors.NewBadRequestError("emptyname")
	}

	if r.Harga <= 0 {
		return errors.NewBadRequestError("emptyprice")
	}

	if r.Kuantitas <= 0 {
		return errors.NewBadRequestError("emptyquantity")
	}

	if r.Deskripsi == "" {
		return errors.NewBadRequestError("emptydescription")
	}

	return nil
}
