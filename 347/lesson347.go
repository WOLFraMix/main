package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// CompressFile читает файл по пути inputPath, сжимает его содержимое
// с помощью RLE-алгоритма и записывает результат в файл по пути outputPath.
// Возвращает ошибку, если произошла проблема при чтении или записи.
func CompressFile(inputPath, outputPath string) error {
	if inputPath == "" {
		return errors.New("input path is empty")
	}
	if outputPath == "" {
		return errors.New("output path is empty")
	}

	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	if len(data) == 0 {
		return os.WriteFile(outputPath, []byte{}, 0644)
	}

	result := make([]byte, 0, len(data)*2)
	runes := []rune(string(data))
	count := 1
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			count++
			if count > 9 {
				result = append(result, fmt.Sprintf("%d%c", 9, runes[i-1])...)
				count -= 9
			}
		} else {
			result = append(result, fmt.Sprintf("%d%c", count, runes[i-1])...)
			count = 1
		}
	}
	result = append(result, fmt.Sprintf("%d%c", count, runes[len(runes)-1])...)

	err = os.WriteFile(outputPath, result, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// Функция main от вас будет скрыта при проверке на сайте
func main() {
	if !test1() || !test2() || !test3() || !test4() || !test5() || !test6() || !test7() || !test8() ||
		!test9() || !test10() || !test11() || !test12() || !test13() {
		os.Exit(1)
	}

	fmt.Println("Все тесты успешно пройдены!")
}

// Обычный случай, базовая проверка
func test1() bool {
	input := "test_input_1.txt"
	output := "test_output_1.txt"

	os.WriteFile(input, []byte("AAABBC"), 0644)
	defer os.Remove(input)
	defer os.Remove(output)

	err := CompressFile(input, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 1: %s\n", err.Error())
		return false
	}

	result, _ := os.ReadFile(output)
	expected := "3A2B1C"
	if string(result) != expected {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 1: Ожидалось: %s, получили: %s\n", expected, string(result))
		return false
	}

	return true
}

// Проверяем, что пустой файл не ломает код
func test2() bool {
	input := "test_input_2.txt"
	output := "test_output_2.txt"
	if err := os.WriteFile(input, []byte{}, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 2: не удалось создать входной файл: %s\n", err.Error())
		return false
	}
	defer os.Remove(input)
	defer os.Remove(output)
	err := CompressFile(input, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 2: %s\n", err.Error())
		return false
	}
	result, err := os.ReadFile(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 2: выходной файл не создан: %s\n", err.Error())
		return false
	}
	if len(result) != 0 {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 2: Ожидался пустой файл, получили: %s\n", string(result))
		return false
	}
	return true
}

// Кириллица (UTF-8)
func test3() bool {
	input := "test_input_3.txt"
	output := "test_output_3.txt"

	os.WriteFile(input, []byte("Привет"), 0644)
	defer os.Remove(input)
	defer os.Remove(output)

	err := CompressFile(input, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 3: %s\n", err.Error())
		return false
	}

	result, _ := os.ReadFile(output)
	expected := "1П1р1и1в1е1т"
	if string(result) != expected {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 3: Ожидалось: %s, получили: %s\n", expected, string(result))
		return false
	}

	return true
}

// Если в файле цифры
func test4() bool {
	input := "test_input_4.txt"
	output := "test_output_4.txt"

	os.WriteFile(input, []byte("112233"), 0644)
	defer os.Remove(input)
	defer os.Remove(output)

	err := CompressFile(input, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 4: %s\n", err.Error())
		return false
	}

	result, _ := os.ReadFile(output)
	expected := "212223"
	if string(result) != expected {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 4: Ожидалось: %s, получили: %s\n", expected, string(result))
		return false
	}

	return true
}

// Файл не существует, ловим ошибку
func test5() bool {
	output := "test_output_5.txt"
	defer os.Remove(output)

	err := CompressFile("non_existent_file.txt", output)
	if err == nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 5: Ожидалась ошибка для несуществующего файла\n")
		return false
	}

	if !errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 5: Ожидалась ошибка IsNotExist, получили: %s\n", err.Error())
		return false
	}

	return true
}

// Одиночные символы без повторов
func test6() bool {
	input := "test_input_6.txt"
	output := "test_output_6.txt"

	os.WriteFile(input, []byte("ABC"), 0644)
	defer os.Remove(input)
	defer os.Remove(output)

	err := CompressFile(input, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 6: %s\n", err.Error())
		return false
	}

	result, _ := os.ReadFile(output)
	expected := "1A1B1C"
	if string(result) != expected {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 6: Ожидалось: %s, получили: %s\n", expected, string(result))
		return false
	}

	return true
}

// Разный регистр букв
func test7() bool {
	input := "test_input_7.txt"
	output := "test_output_7.txt"

	os.WriteFile(input, []byte("aAbBcC"), 0644)
	defer os.Remove(input)
	defer os.Remove(output)

	err := CompressFile(input, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 7: %s\n", err.Error())
		return false
	}

	result, _ := os.ReadFile(output)
	expected := "1a1A1b1B1c1C"
	if string(result) != expected {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 7: Ожидалось: %s, получили: %s\n", expected, string(result))
		return false
	}

	return true
}

// Длинные повторения
func test8() bool {
	input := "test_input_8.txt"
	output := "test_output_8.txt"

	os.WriteFile(input, []byte("AAAAABBBBBCCCCC"), 0644)
	defer os.Remove(input)
	defer os.Remove(output)

	err := CompressFile(input, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 8: %s\n", err.Error())
		return false
	}

	result, _ := os.ReadFile(output)
	expected := "5A5B5C"
	if string(result) != expected {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 8: Ожидалось: %s, получили: %s\n", expected, string(result))
		return false
	}

	return true
}

// Повторяющиеся символы через другие, тут идет проверка сброса счетчика.
func test9() bool {
	input := "test_input_9.txt"
	output := "test_output_9.txt"

	os.WriteFile(input, []byte("abracadabra"), 0644)
	defer os.Remove(input)
	defer os.Remove(output)

	err := CompressFile(input, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 9: %s\n", err.Error())
		return false
	}

	result, _ := os.ReadFile(output)
	expected := "1a1b1r1a1c1a1d1a1b1r1a"
	if string(result) != expected {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 9: Ожидалось: %s, получили: %s\n", expected, string(result))
		return false
	}

	return true
}

// Выходной путь - это существующий каталог, запись должна закончиться ошибкой
func test10() bool {
	input := "test_input_10.txt"
	os.WriteFile(input, []byte("AAABBC"), 0644)
	defer os.Remove(input)

	output, err := os.MkdirTemp("", "test_output_dir_10")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 10: не удалось создать каталог: %s\n", err.Error())
		return false
	}
	defer os.Remove(output) // каталог пуст, Remove достаточно

	if err := CompressFile(input, output); err == nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 10: ожидалась ошибка записи в каталог\n")
		return false
	}
	return true
}

// Выходной файл лежит в несуществующем каталоге
func test11() bool {
	input := "test_input_11.txt"
	os.WriteFile(input, []byte("AAABBC"), 0644)
	defer os.Remove(input)

	err := CompressFile(input, "no_such_dir_11/test_output.txt")
	if err == nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 11: ожидалась ошибка записи\n")
		return false
	}
	if !errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 11: ожидалась ErrNotExist, получили: %s\n", err.Error())
		return false
	}
	return true
}

// Серия длиннее 9 символов должна дробиться на части. Счетчик - это всегда одна цифра.
func test12() bool {
	input := "test_input_12.txt"
	output := "test_output_12.txt"

	content := strings.Repeat("A", 101) // 101 буква A
	if err := os.WriteFile(input, []byte(content), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 12: не удалось создать входной файл: %s\n", err.Error())
		return false
	}
	defer os.Remove(input)
	defer os.Remove(output)

	if err := CompressFile(input, output); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 12: %s\n", err.Error())
		return false
	}

	result, _ := os.ReadFile(output)
	expected := strings.Repeat("9A", 11) + "2A" // 9*11 + 2 = 101
	if string(result) != expected {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 12: Ожидалось: %s, получили: %s\n", expected, string(result))
		return false
	}
	return true
}

// Длинная серия цифр режется так же, как буквы
func test13() bool {
	input := "test_input_13.txt"
	output := "test_output_13.txt"

	os.WriteFile(input, []byte(strings.Repeat("7", 23)+"x"), 0644)
	defer os.Remove(input)
	defer os.Remove(output)

	if err := CompressFile(input, output); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 13: %s\n", err.Error())
		return false
	}

	result, _ := os.ReadFile(output)
	expected := "97" + "97" + "57" + "1x" // 23 = 9 + 9 + 5
	if string(result) != expected {
		fmt.Fprintf(os.Stderr, "Ошибка в тесте 13: Ожидалось: %s, получили: %s\n", expected, string(result))
		return false
	}
	return true
}
