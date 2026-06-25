package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "    Hello,   World ! "
	strNormal := "Привет, мир!"
	fmt.Println(str)
	fmt.Println(strings.TrimSpace(str))                       // удаляет "лишние" пробелы только слева и справа
	fmt.Println(strings.ToUpper(str))                         // переводит в верхний регистр
	fmt.Println(strings.ToLower(strNormal))                   // переводит в нижний регистр
	fmt.Println(strings.Split(strNormal, " "))                // разделяет строку на массив строк по разделителю (пробел)
	fmt.Println(strings.Contains(strNormal, "!"))             // проверяет, содержится ли "!" в строке
	fmt.Println(strings.HasPrefix(strNormal, "Пр"))           // проверяет, начинается ли строка с "Пр"
	fmt.Println(strings.HasSuffix(strNormal, "мир!"))         // проверяет, заканчивается ли строка на "мир!"
	fmt.Println(strings.EqualFold(strNormal, "ПриВеТ, МИР!")) // проверяет, равны ли строки без учета регистра
}
