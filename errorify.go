package errorify

import (
	"fmt"
)

const (
	// RED color
	RED = "\033[0;31m"
	// GREEN color
	GREEN = "\033[0;32m"
	// YELLOW color
	YELLOW = "\033[1;33m"
	// ENDLINE end of text line (close color tag)
	ENDLINE = "\033[0m"
)

var (
	// Language ...
	// TODO: create description
	Language string
	message  string
	filename string
)

// CreateFile ...
// TODO: create description
func CreateFile(err error, params ...string) {
	if err != nil {
		// Check if params have filename
		if len(params) == 1 {
			filename = "'" + params[0] + "' "
		}

		// Switching to lang
		switch Language {
		case "ru_RU": // Russian
			message = RED + "✘ Ошибка: файл " + filename + "не был создан!" + ENDLINE
		case "es_ES": // Español
			message = RED + "✘ Error: ¡No se creó el archivo " + filename + "!" + ENDLINE
		default:
			message = RED + "✘ Error: file " + filename + "not created!" + ENDLINE
		}

		// Print result
		fmt.Println(message)
		return
	}
}

// Simple ...
// TODO: create description
func Simple(err error) {
	if err != nil {
		// Switching to lang
		switch Language {
		case "ru_RU": // Russian
			message = RED + "✘ Error: " + err.Error() + ENDLINE
		case "es_ES": // Español
			message = RED + "✘ Error: " + err.Error() + ENDLINE
		default:
			message = RED + "✘ Error: " + err.Error() + ENDLINE
		}

		// Print result
		fmt.Println(message)
		return
	}
}
