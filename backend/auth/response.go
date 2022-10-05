package auth

type DataResponse struct {
	Email string `json:"email"`
	//ID    uint   `json:"id"`
	Token string `json:"token"`
}
