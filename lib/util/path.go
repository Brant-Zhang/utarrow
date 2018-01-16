package util

import (
	"path/filepath"
	"runtime"
	"strings"
)

func GetCurrentFileName() string {
	_, file, _, _ := runtime.Caller(1)
	f := strings.Split(file, "/")
	fileName := strings.Replace(f[len(f)-1], ".go", "", 1)
	return fileName
}

func GetCurrentFilePath() string {
	_, file, _, _ := runtime.Caller(1)
	f := []string{"/"}
	_f := strings.Split(file, "/")
	f = append(f, _f[:len(_f)-1]...)
	return filepath.Join(f...)
}
