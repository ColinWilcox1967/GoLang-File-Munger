package fileutils

import (
	"os"
	"errors"
	"io/ioutil"
	"strings"
	"path/filepath"
)

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	 }

	 return true
}

func ReadFileAsBytes(filePath string) (bool, []byte) {
	v, err := ioutil.ReadFile(filePath)  
	if err != nil {
		return false, []byte{}
	}

	return true, v
}

func GetFileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
