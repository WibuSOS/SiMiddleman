package localizator

import (
	"fmt"

	// "github.com/WibuSOS/sinarmas/backend/utils/localization"
	"github.com/gin-gonic/gin"
	language "github.com/moemoe89/go-localization"
)

type Handler struct {
	Service *language.Config
}

func NewHandler() (*Handler, error) {
	return nil, fmt.Errorf("error")
	// lang, err := localization.Initialize()
	// if err != nil {
	// 	return nil,err
	// }
	// return &Handler{Service: lang},nil
}

func (h *Handler) PassLocalizator(c *gin.Context) {
}

func (h *Handler) GetMessage() {
}
