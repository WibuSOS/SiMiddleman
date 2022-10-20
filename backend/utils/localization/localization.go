package localization

import (
	"encoding/json"
	"io/fs"
	"os"

	"github.com/WibuSOS/sinarmas/backend/helpers"
)

type data struct {
	En map[string]string `json:"en"`
	Id map[string]string `json:"id"`
}

func readDirectory(dir string) []fs.FileInfo {
	f, _ := os.Open(dir)
	files, _ := f.Readdir(0)

	return files
}

func readJSON(file string) map[string]interface{} {
	content, _ := os.ReadFile(file)
	var payload map[string]interface{}
	json.Unmarshal(content, &payload)

	return payload
}

func collectData() data {
	en := map[string]string{}
	id := map[string]string{}

	rootPath := helpers.GetRootPath()

	middlewaresPath := rootPath + "/middlewares"
	dirMiddleware := readDirectory(middlewaresPath)
	for _, v := range dirMiddleware {
		if v.IsDir() && v.Name() != "localizator" {
			subPath := middlewaresPath + "/" + v.Name()
			files := readDirectory(subPath)
			for _, v := range files {
				if v.Name() == "language.json" {
					data := readJSON(subPath + "/" + v.Name())
					for key, value := range data {
						rec := value.(map[string]interface{})
						if key == "en" {
							for k, val := range rec {
								en[k] = val.(string)
							}
						} else if key == "id" {
							for k, val := range rec {
								id[k] = val.(string)
							}
						}
					}
				}
			}
		}
	}

	controllersPath := rootPath + "/controllers"
	dirController := readDirectory(controllersPath)
	for _, v := range dirController {
		if v.IsDir() {
			subPath := controllersPath + "/" + v.Name()
			files := readDirectory(subPath)
			for _, v := range files {
				if v.Name() == "language.json" {
					data := readJSON(subPath + "/" + v.Name())
					for key, value := range data {
						rec := value.(map[string]interface{})
						if key == "en" {
							for k, val := range rec {
								en[k] = val.(string)
							}
						} else if key == "id" {
							for k, val := range rec {
								id[k] = val.(string)
							}
						}
					}
				}
			}
		}
	}

	allData := data{
		En: en,
		Id: id,
	}

	return allData
}

func WriteJSON() {
	rootPath := helpers.GetRootPath()

	data := collectData()

	content, _ := json.Marshal(data)
	os.WriteFile(rootPath+"/middlewares/localizator/language.json", content, 0644)
}
