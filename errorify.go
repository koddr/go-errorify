package errorify

import (
	"fmt"
)

const (
	// RED ...
	RED = "\033[0;31m"
	// GREEN ...
	GREEN = "\033[0;32m"
	// YELLOW ...
	YELLOW = "\033[1;33m"
	// ENDLINE ...
	ENDLINE = "\033[0m"
)

var message error

// CreateFile ...
func CreateFile(err error, name string) {
	if err != nil {
		message = fmt.Errorf(RED+"✘ Error: file '%v' not created!"+ENDLINE, name)
		fmt.Println(message)
		return
	}
}

// Simple ...
func Simple(err error) {
	if err != nil {
		message = fmt.Errorf(RED+"✘ Error: %v"+ENDLINE, err)
		fmt.Println(message)
		return
	}
}
