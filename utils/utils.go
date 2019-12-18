package utils

import (
	"os"
	"path/filepath"
)

func GetRoot() string {
	root,_ := filepath.Abs(filepath.Dir(os.Args[0]))
	return string(root)
}


func GetPath(path string) string {
	return GetRoot()+"/"+path
}