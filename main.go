package modtest

import (
	"fmt"
	"os"
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

	// from Executable Directory
	ex, _ := os.Executable()
	fmt.Println("Executable DIR:", filepath.Dir(ex))

	// Current working directory
	dir, _ := os.Getwd()
	fmt.Println("CWD:", dir)

	// Relative on runtime DIR:
	_, b, _, _ := runtime.Caller(0)
	d1 := path.Join(path.Dir(b))
	fmt.Println("Relative", d1)

}
