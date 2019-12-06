package main

import (
	"github.com/koddr/go-errorify"
	"os"
)

func init() {
	// init with ru_RU language
	errorify.Language = "ru_RU"
}

func main() {
	// set 3 files
	fileNameFirst := "./folder/file_1.txt"
	fileNameSecond := "./folder/file_2.txt"
	fileNameThird := "./folder/file_3.txt"

	// first file
	// without any params
	a, err := os.Create(fileNameFirst)
	errorify.CreateFile(err) // ✘ Ошибка: файл не был создан!
	defer a.Close()

	// second file
	// with filename param
	b, err := os.Create(fileNameSecond)
	errorify.CreateFile(err, fileNameSecond) // ✘ Ошибка: файл './folder/file_2.txt' не был создан!
	defer b.Close()

	// switch language to es_ES
	errorify.Language = "es_ES"

	// third file
	// with filename param
	c, err := os.Create(fileNameThird)
	errorify.CreateFile(err, fileNameThird) // ✘ Error: ¡No se creó el archivo './folder/file_3.txt' !
	defer c.Close()
}
