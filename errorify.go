package errorify

import (
	"fmt"
)

const (
	red    = "\033[0;31m" // red color
	yellow = "\033[1;33m" // yellow color
	end    = "\033[0m"    // end of text line (close color tag)
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
			message = red + "✘ Ошибка: файл " + filename + "не был создан!" + end
		case "es_ES": // Español
			message = red + "✘ Error: ¡No se creó el archivo " + filename + "!" + end
		default:
			message = red + "✘ Error: file " + filename + "not created!" + end
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
			message = red + "✘ Error: " + err.Error() + end
		case "es_ES": // Español
			message = red + "✘ Error: " + err.Error() + end
		default:
			message = red + "✘ Error: " + err.Error() + end
		}

		// Print result
		fmt.Println(message)
		return
	}
}
