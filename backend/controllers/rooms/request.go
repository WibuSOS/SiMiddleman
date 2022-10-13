package rooms

import (
	"github.com/WibuSOS/sinarmas/backend/controllers/product"
	"github.com/WibuSOS/sinarmas/backend/utils/errors"
)

type DataRequest struct {
	PenjualID uint                 `json:"id" binding:"required"`
	Product   *product.DataRequest `json:"product" binding:"required"`
}

func (req *DataRequest) ValidateReq() *errors.RestError {
	if err := req.Product.ValidateReq(); err != nil {
		return err
	}

	return nil
}
