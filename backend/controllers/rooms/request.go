package rooms

import (
	"github.com/WibuSOS/sinarmas/utils/errors"
)

type DataRequest struct {
	PenjualID uint                `json:"id"`
	Produk    product.DataRequest `json:"produk"`
}

func (req *DataRequest) ValidateReq() *errors.RestError {
	if req.PenjualID == 0 {
		return errors.NewBadRequestError("oops... there is something wrong")
	}

	if err := req.Produk.ValidateReq(); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}
