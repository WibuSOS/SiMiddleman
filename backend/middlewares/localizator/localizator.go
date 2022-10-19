package localizator

import (
	"fmt"
	"os"

	"github.com/WibuSOS/sinarmas/backend/helpers"
	"github.com/gin-gonic/gin"
	language "github.com/moemoe89/go-localization"
)

type Handler struct {
	Service *language.Config
}

func NewHandler() (*Handler, error) {
	lang, err := initialize()
	if err != nil {
		return nil, err
	}

	return &Handler{Service: lang}, nil
}

func initialize() (*language.Config, error) {
	dir := helpers.GetRootPath() + os.Getenv("LOCALIZATOR_PATH")
	path := fmt.Sprintf("%s/language.json", dir)

	cfg := language.New()
	cfg.BindPath(path)
	cfg.BindMainLocale("en")

	lang, err := cfg.Init()
	if err != nil {
		return nil, err
	}

	return lang, nil
}

func (h *Handler) PassLocalizator(c *gin.Context) {
	c.Set("localizator", h.Service)
	c.Next()
}
