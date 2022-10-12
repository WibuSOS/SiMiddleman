package transaction

type ResponsePaymentInfo struct {
	Total  uint   `json:"total"`
	Status string `json:"status"`
}
