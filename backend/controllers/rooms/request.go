package rooms

import (
	"github.com/WibuSOS/sinarmas/controllers/product"
	"github.com/WibuSOS/sinarmas/utils/errors"
)

type DataRequest struct {
	PenjualID uint                 `json:"id"`
	Product   *product.DataRequest `json:"product"`
}

func (req *DataRequest) ValidateReq() *errors.RestError {
	if req.PenjualID == 0 {
		return errors.NewBadRequestError("oops... there is something wrong")
	}

	if err := req.Product.ValidateReq(); err != nil {
		return err
	}

	return nil
}
