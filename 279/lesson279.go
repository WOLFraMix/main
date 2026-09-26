package main

import "fmt"

// Интерфейс Killer с методом Confess() string
type Killer interface {
	Confess() string
}

// Расследование
func Investigation(suspects []Suspect) {
	fmt.Println("НАЧАЛО ДОПРОСА")
	defer fmt.Println("ДОПРОС ОКОНЧЕН")
	for _, sus := range suspects {
		fmt.Println(sus.Alibi())
	}
	for _, sus := range suspects {
		killer, ok := sus.(Killer)
		if ok {
			fmt.Println("Дуайт: АГА! ПОПАЛСЯ!")
			fmt.Println(killer.Confess())
			return
		}
	}
	fmt.Println("Майкл: Больше мафиози я ненавижу только врунов. Вот бы мафиози поубивали всех врунов.")
}

// ------------------------
// Код ниже не трогаем, некоторые детали (методы структур) от вас скрыты

// Suspect базовый интерфейс подозреваемого
type Suspect interface {
	Alibi() string
}

// Подозреваемые
type Michael struct{}
type Dwight struct{}
type Darryl struct{}
type Creed struct{}
type Toby struct{}

func main() {
	// Дуайт собирает всех в переговорке
	suspects := []Suspect{
		Michael{},
		Dwight{},
		Darryl{},
		Creed{},
		Toby{},
	}

	Investigation(suspects)
}
