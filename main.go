package modtest

import (
	"fmt"
	"path/filepath"
)

func Print() {
	fmt.Println("v1.0.1")
}

func PrintPath() {
	abs, err := filepath.Abs("./data/")
	if err == nil {
		fmt.Println("Absolute:", abs)
	}
}
