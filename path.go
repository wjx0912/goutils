package goutils

import (
	"path/filepath"
	"os"
	"strings"
)

func GetModulePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1) + "/"
}

func SetWorkPath2Module() bool {
	path := GetModulePath()
	path = "D:/git/distribute_gtp/client_gtpserver/bin"	// TODO: remove this debug line
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
