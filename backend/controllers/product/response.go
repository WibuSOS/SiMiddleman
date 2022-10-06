package product

type DataResponse struct {
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
	Harga     uint   `json:"harga"`
	Kuantitas uint   `json:"kuantitas"`
}
