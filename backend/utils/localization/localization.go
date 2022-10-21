package localization

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"

	"github.com/WibuSOS/sinarmas/backend/helpers"
)

type data struct {
	En map[string]string `json:"en"`
	Id map[string]string `json:"id"`
}

func check(e error) {
	if e != nil {
		log.Println(e.Error())
		panic(e)
	}
}

func readDirectory(dir string) []fs.DirEntry {
	// rootPath := helpers.GetRootPath()
	// root := strings.ReplaceAll(rootPath, `/`, `\`)

	// dir = filepath.Join(root, filepath.Clean(dir))
	// if !strings.HasPrefix(dir, root) {
	// 	panic(fmt.Errorf("unsafe input"))
	// }
	f, err := os.ReadDir(dir)
	check(err)

	return f
}

func readJSON(file string) map[string]interface{} {
	// rootPath := helpers.GetRootPath()
	// root := strings.ReplaceAll(rootPath, `/`, `\`)

	// file = filepath.Join(root, filepath.Clean(file))
	// if !strings.HasPrefix(file, root) {
	// 	panic(fmt.Errorf("unsafe input"))
	// }
	content, err := os.ReadFile(file)
	check(err)

	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	check(err)

	return payload
}

func collectData() data {
	en := map[string]string{}
	id := map[string]string{}

	rootPath := helpers.GetRootPath()

	middlewaresPath := rootPath + "/middlewares"
	dirMiddleware := readDirectory(middlewaresPath)
	for _, v := range dirMiddleware {
		if !v.IsDir() || v.Name() == "localizator" {
			continue
		}

		subPath := middlewaresPath + "/" + v.Name()
		files := readDirectory(subPath)
		for _, v := range files {
			if v.Name() != "language.json" {
				continue
			}

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
			break
		}
	}

	controllersPath := rootPath + "/controllers"
	dirController := readDirectory(controllersPath)
	for _, v := range dirController {
		if !v.IsDir() {
			continue
		}

		subPath := controllersPath + "/" + v.Name()
		files := readDirectory(subPath)
		for _, v := range files {
			if v.Name() != "language.json" {
				continue
			}

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
			break
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
	err := os.WriteFile(rootPath+os.Getenv("LOCALIZATOR_PATH")+"/language.json", content, 0600)
	check(err)
}
