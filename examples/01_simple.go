package main

import (
	"github.com/koddr/go-errorify"
	"os"
)

func main() {
	// Set 1 files
	fileNameFirst := "./folder/file_1.txt"

	// First file
	a, err := os.Create(fileNameFirst)
	errorify.Simple(err) // ✘ Error: open ./folder/file_1.txt: no such file or directory
	defer a.Close()
}
