package goutils

import (
	"path/filepath"
	"os"
	"strings"
)

func GetModuleDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1) + "/"
}

func SetWorkDirectory(path string) bool {
	err := os.Chdir(path)
	if(err != nil) {
		return false
	} else {
		return true
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
