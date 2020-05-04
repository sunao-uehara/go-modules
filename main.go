package modtest

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
)

func Print() {
	fmt.Println("v1.0.1")
}

func PrintPath() {
	abs, err := filepath.Abs("./data/")
	if err == nil {
		fmt.Println("Absolute:", abs)
	}

	_, filename, _, ok := runtime.Caller(1)
	if ok {
		fmt.Println("filename:", path.Dir(filename))
	}
}
