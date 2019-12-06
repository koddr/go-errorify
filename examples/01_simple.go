package main

import (
	"github.com/koddr/go-errorify"
	"os"
)

func main() {
	// set 1 files
	fileNameFirst := "./folder/file_1.txt"

	// first file
	a, err := os.Create(fileNameFirst)
	errorify.Simple(err) // ✘ Error: open ./folder/file_1.txt: no such file or directory
	defer a.Close()
}
