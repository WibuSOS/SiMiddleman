package helpers

import (
	"path/filepath"
	"runtime"
	"strings"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func GetRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	// b, _ := os.Getwd()
	// Root folder of this project
	root := strings.ReplaceAll(filepath.Join(filepath.Dir(b), ""), `\`, `/`)
	// root := strings.ReplaceAll(b, `\`, `/`)

	return root
}
