package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Запускаем обход текущей директории "."
	PrintAllFiles(".")
}

// PrintAllFiles рекурсивно выводит все элементы (файлы и папки) в указанной директории
func PrintAllFiles(path string) {
	// Получаем список элементов в папке (файлы и директории)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to get list of files in %q: %v\n", path, err)
		return
	}

	// Проходим по каждому элементу
	for _, f := range files {
		// Собираем полный путь к элементу с учётом разделителей ОС
		filename := filepath.Join(path, f.Name())

		// Выводим путь
		fmt.Println(filename)

		// Если это директория, рекурсивно обходим её
		if f.IsDir() {
			PrintAllFiles(filename)
		}
	}
}
