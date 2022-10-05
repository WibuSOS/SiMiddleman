package rooms

import (
	"github.com/WibuSOS/sinarmas/product"
	"github.com/WibuSOS/sinarmas/utils/errors"
)

type DataRequest struct {
	PenjualID uint                `json:"id"`
	Product   *DataRequestProduct `json:"produk"`
}

type DataRequestProduct struct {
	Nama      string `json:"nama"`
	Harga     int    `json:"harga"`
	Kuantitas int    `json:"kuantitas"`
	Deskripsi string `json:"deskripsi"`
}

func (req *DataRequest) ValidateReq() *errors.RestError {
	if req.PenjualID == 0 {
		return errors.NewBadRequestError("oops... there is something wrong")
	}

	if err := req.Produk.ValidateReq(); err != nil {
		return err
	}

	return nil
}

func (p *DataRequestProduct) ValidateReq() error {
	if p.Nama == "" {
		return fmt.Errorf("nama tidak memenuhi syarat")
	}

	if p.Deskripsi == "" {
		return fmt.Errorf("deskripsi tidak memenuhi syarat")
	}

	if p.Harga == 0 {
		return fmt.Errorf("harga tidak memenuhi syarat")
	}

	if p.Kuantitas == 0 {
		return fmt.Errorf("kuantitas tidak memenuhi syarat")
	}

	return nil
}
