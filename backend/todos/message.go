package todos

type DataRequest struct {
	Task string `json:"task" binding:"required"`
}
