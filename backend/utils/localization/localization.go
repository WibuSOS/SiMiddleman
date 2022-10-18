package localization

import (
	"fmt"
	"os"
	"strings"

	language "github.com/moemoe89/go-localization"
)

func initialize() (*language.Config, error) {
	dir, _ := os.Getwd()
	strDir := strings.ReplaceAll(dir, `\`, `/`)
	path := fmt.Sprintf("%s/utils/localization/language.json", strDir)

	cfg := language.New()
	cfg.BindPath(path)
	cfg.BindMainLocale("en")

	lang, err := cfg.Init()
	if err != nil {
		return nil, err
	}

	return lang, nil
}

func GetMessage(langReq, id string) string {
	lang, err := initialize()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return lang.Lookup(langReq, id)
}
