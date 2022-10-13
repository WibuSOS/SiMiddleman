package transaction

import (
	"github.com/WibuSOS/sinarmas/backend/helpers"
	"github.com/WibuSOS/sinarmas/backend/utils/errors"
)

type RequestUpdateStatus struct {
	Status string `json:"status"`
}

func (r *RequestUpdateStatus) Validation() *errors.RestError {
	statusArr := []string{"mulai transaksi", "barang dibayar", "barang dikirim", "konfirmasi barang sampai"}

	if r.Status == "" || !helpers.Contains(statusArr, r.Status) {
		return errors.NewBadRequestError("Invalid Status")
	}

	return nil
}
