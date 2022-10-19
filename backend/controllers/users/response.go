package users

type DataResponse struct {
	Nama  string `json:"nama" binding:"required"`
	NoHp  string `json:"noHp" binding:"required"`
	Email string `json:"email" binding:"required"`
	NoRek string `json:"noRek" binding:"required"`
}
