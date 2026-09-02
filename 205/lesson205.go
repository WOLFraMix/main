package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type VillageElement interface {
	Update()
	FlushInfo() string
}

type Village struct {
	Elements []VillageElement
}

type Resident struct {
	Name    string
	Age     int
	Married bool
	Alive   bool
	Events  []string
}

func (r *Resident) AddYear() {
	r.Age++
}

func (r *Resident) ChangeMarried() {
	if r.Married == false {
		r.Married = true
		r.Events = append(r.Events, "Свадьба! Ура.")
	} else {
		r.Events = append(r.Events, "Развод. Больше не в браке...")
	}
}

func (r *Resident) Death() {
	r.Alive = false
	r.Events = append(r.Events, fmt.Sprintf("Покидает мир на %d году жизни", r.Age))
}

func (r *Resident) Update() {
	if !r.Alive {
		return
	}
	r.AddYear()
	wasMarried := r.Married // состояние ДО любых изменений в этом году

	if rand.IntN(100) < 10 {
		r.ChangeMarried()
	}
	if rand.IntN(100) < 20 {
		r.Events = append(r.Events, "Нашлась новая работа!")
	}
	if rand.IntN(100) < 15 {
		r.Events = append(r.Events, "Куплена новая машина.")
	}
	if rand.IntN(100) < 15 {
		r.Events = append(r.Events, "Потерялся телефон(((.")
	}
	if rand.IntN(100) < 20 {
		r.Events = append(r.Events, "Появилось новое хобби.")
	}
	if wasMarried && r.Married && rand.IntN(100) < 50 {
		r.Events = append(r.Events, "Был хороший отпуск. Все отдохнули!")
	}
	if rand.IntN(100) < 5 {
		r.Death()
	}
}

func (r *Resident) FlushInfo() string {
	info := fmt.Sprintf("Житель %s умер в возрасте %d лет.\n", r.Name, r.Age)
	if r.Alive {
		married := "свободный"
		if r.Married {
			married = "в браке"
		}
		events := "нету"
		if len(r.Events) > 0 {
			events = strings.Join(r.Events, "\n")
		}
		info = fmt.Sprintf("Житель %s (возраст %d лет), статус %s.\nСобытия: %s\n", r.Name, r.Age, married, events)
	}
	r.Events = []string{}
	return info
}

func (v *Village) AddElement(e VillageElement) {
	v.Elements = append(v.Elements, e)
}

func (v *Village) UpdateAll() {
	for _, e := range v.Elements {
		e.Update()
	}
}

func (v Village) ShowAllInfo() string {
	info := ""
	for _, e := range v.Elements {
		info += e.FlushInfo()
	}
	return info
}

type Animal struct {
	Name   string
	Age    int
	Type   string
	Alive  bool
	Events []string
}

func (a *Animal) AddYear() {
	a.Age++
}

func (a *Animal) Death() {
	a.Alive = false
	a.Events = append(a.Events, fmt.Sprintf("Бедная %s. Покидает наш мир на %d году жизни.\n", a.Type, a.Age))
}

func (a *Animal) Update() {
	if !a.Alive {
		return
	}
	a.AddYear()
	if rand.IntN(100) < 15 {
		a.Events = append(a.Events, "Погонялась за хвостом!")
	}
	if rand.IntN(100) < 5 {
		a.Events = append(a.Events, "Сломала лапу.")
	}
	if rand.IntN(100) < 15 {
		a.Events = append(a.Events, "Поймала добычу!")
	}
	if rand.IntN(100) < 10 {
		a.Events = append(a.Events, "Пропала и снова нашлась!")
	}
	if rand.IntN(100) < 10 {
		a.Death()
	}
}

func (a *Animal) FlushInfo() string {
	info := fmt.Sprintf("Животное: %s %s - умерло в возрасте %d лет.\n", a.Type, a.Name, a.Age)
	if a.Alive {
		events := "нету"
		if len(a.Events) > 0 {
			events = strings.Join(a.Events, "\n")
		}
		info = fmt.Sprintf("Животное %s (%s, возраст %d лет).\nСобытия: %s\n", a.Name, a.Type, a.Age, events)
	}
	a.Events = []string{}
	return info
}

func main() {
	village := Village{}

	// Создаем жителей деревни
	resident1 := &Resident{Name: "Алиса", Age: 30, Married: false, Alive: true, Events: []string{}}
	resident2 := &Resident{Name: "Борис", Age: 40, Married: true, Alive: true, Events: []string{}}

	// Создаем животных
	animal1 := &Animal{Name: "Бобик", Age: 5, Type: "собака", Alive: true, Events: []string{}}
	animal2 := &Animal{Name: "Мурка", Age: 3, Type: "кошка", Alive: true, Events: []string{}}

	// Добавляем элементы в деревню
	village.AddElement(resident1)
	village.AddElement(resident2)
	village.AddElement(animal1)
	village.AddElement(animal2)

	// Симуляция обновления деревни на несколько лет
	for i := 0; i < 5; i++ {
		fmt.Printf("Год %d:\n", i+1)
		village.UpdateAll()
		fmt.Println(village.ShowAllInfo())
	}
}
