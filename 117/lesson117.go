package main

import "fmt"

func main() {
	m := make(map[string]string) // создаём map
	m["foo"] = "bar"             // заполняем парами «ключ-значение»
	m["ping"] = "pong"
	fmt.Println(m)

	// var m map[KeyType]ValueType

	// Переменные типа map инициализируются с помощью функции make()

	type MyMap map[string]string

	var m1 MyMap
	m1 = make(MyMap, 5)
	m1["high"] = "five"
	fmt.Println(m1)

	x := 5
	y := x
	x++
	// x станет равен 6
	// y останется равен 5
	fmt.Println(x, y)

	MyMap1 := make(MyMap)
	MyMap2 := MyMap1
	MyMap1["foo"] = "bar"
	// в MyMap2 тоже появится пара с ключом foo и значением bar
	// если поменяем значение в MyMap2,
	MyMap2["foo"] = "bazz" // по этому ключу изменится значение
	// то изменится значение и в MyMap1
	fmt.Println(MyMap1)
	// если добавим другой ключ, то добавится и новое значение
	MyMap1["zoo"] = "park"
	fmt.Println(MyMap2)

	MyStringMap := map[string]string{"first": "первый", "second": "второй"}
	// в данном случае композитный литерал создаёт map без использования функции make
	// и уже с инициализированными парами «ключ-значение»
	fmt.Println(MyStringMap)

	var m2 map[int]int
	m3 := map[int]int{1: 10, 2: 20, 3: 30}
	fmt.Println(len(m2), len(m3))

	mmm := map[int]string{1: "first"}
	v, ok := mmm[1]
	fmt.Println(v, ok)
	delete(mmm, 2) // ошибки не будет
	delete(mmm, 1)
	v, ok = mmm[1]
	fmt.Println(v, ok)

	mr := make(map[string]string)
	mr["foo"] = "bar"
	mr["bazz"] = "yup"
	for k, val := range mr {
		// k будет перебирать ключи,
		// val — соответствующие этим ключам значения
		fmt.Printf("Ключ %v, имеет значение %v \n", k, val)
	}

	/*
		for k, v := range m {
		m[k] = "here key "+k    // применяем к таблице индексное выражение
		// и модифицируем её прямым доступом к ячейкам
		}
	*/
}
