package main

// Импорты ниже нужны для скрытой проверки (os.Stdin, парсинг JSON) и готовой функции сортировки.
// Вам для решения понадобится только пакет "fmt". Удалять импорты не нужно, добавлять можно, если понадобится.
import (
	"fmt"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

// EffectType описывает тип эффектов, которые может дать ингредиент
type EffectType int

const (
	Strength EffectType = iota
	Agility
	Endurance
	Intelligence
	Luck
)

func (e EffectType) String() string {
	switch e {
	case Strength:
		return "Сила"
	case Agility:
		return "Ловкость"
	case Endurance:
		return "Выносливость"
	case Intelligence:
		return "Интеллект"
	case Luck:
		return "Удача"
	default:
		return "Неизвестно"
	}
}

// Effect представляет собой единичный эффект (например, +5 к силе)
type Effect struct {
	Type  EffectType
	Value int
}

// Ingredient представляет базовый интерфейс для любого ингредиента (даже нейтрального наполнителя).
// Любой ингредиент умеет называться и имеет количество.
type Ingredient interface {
	Name() string
	Amount() int
}

// Beneficial представляет интерфейс для полезных (или вредных) ингредиентов.
// Он встраивает в себя Ingredient, но добавляет метод Effects().
type Beneficial interface {
	Ingredient
	Effects() []Effect
}

// CreateMixture принимает ингредиенты и выводит результат смешивания напрямую в stdout.
func CreateMixture(ingredients ...Ingredient) {
	if len(ingredients) == 0 {
		fmt.Println("Ингредиенты:")
		fmt.Println("Отсутствуют")
		fmt.Println()
		fmt.Println("Эффекты:")
		fmt.Println("Отсутствуют")
		return
	}
	nameCount := make(map[string]int)
	nameC := []string{}
	effectSum := make(map[EffectType]int)
	for _, ingredient := range ingredients {
		_, ok := nameCount[ingredient.Name()]
		if !ok {
			nameC = append(nameC, ingredient.Name())
		}
		nameCount[ingredient.Name()] += ingredient.Amount()
		effects, ok := ingredient.(Beneficial)
		if ok {
			for _, effect := range effects.Effects() {
				effectSum[effect.Type] += effect.Value
			}
		}
	}
	SortIngredientNames(nameC)

	fmt.Println("Ингредиенты:")
	if len(nameC) == 0 {
		fmt.Println("Отсутствуют")
	} else {
		for _, name := range nameC {
			value := nameCount[name]
			fmt.Printf("- %s: %d\n", name, value)
		}
	}
	fmt.Println()
	fmt.Println("Эффекты:")
	if len(effectSum) == 0 {
		fmt.Println("Отсутствуют")
	} else {
		hasEffects := false
		for et := Strength; et <= Luck; et++ {
			value, ok := effectSum[et]
			if !ok || value == 0 {
				continue
			}
			hasEffects = true
			fmt.Printf("- %s: %d\n", et.String(), value)
		}
		if !hasEffects {
			fmt.Println("Отсутствуют")
		}
	}
}

// cl тяжеловесный объект-сортировщик, который нужно, инициализировать только один раз и переиспользовать.
var cl = collate.New(language.Russian)

// SortIngredientNames сортирует слайс строк по правилам русского алфавита (с учетом буквы "ё").
// Обычная сортировка по Unicode ломает правильный порядок, поэтому здесь используется collate.
func SortIngredientNames(names []string) {
	cl.SortStrings(names)
}
