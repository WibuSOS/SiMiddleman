package localizator

import (
	"fmt"
	"os"
	"strings"

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
	dir, _ := os.Getwd()
	strDir := strings.ReplaceAll(dir, `\`, `/`)
	path := fmt.Sprintf("%s/language.json", strDir)

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
	c.Set("localization", h)
}

func (h *Handler) GetMessage(langReq, id string) string {
	return h.Service.Lookup(langReq, id)
}
