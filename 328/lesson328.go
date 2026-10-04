package main

import (
	"fmt"
	"sort"
	"time"
)

type MonsterContract struct {
	ContractID int       `json:"contract_id"`
	Monster    string    `json:"monster"`
	TakenAt    time.Time `json:"taken_at"`
	SlainAt    time.Time `json:"slain_at"`
}

// PrintChronicles принимает слайс контрактов, сортирует их по правилам
// и выводит лог хроник в консоль.
func PrintChronicles(contracts []MonsterContract) {
	if len(contracts) == 0 {
		fmt.Println("Нет данных о контрактах")
		return
	}

	layout := "02.01.2006 15:04"

	// Сортировка: TakenAt → SlainAt → ContractID
	sort.Slice(contracts, func(i, j int) bool {
		cmp1 := contracts[i].TakenAt.Compare(contracts[j].TakenAt)
		if cmp1 != 0 {
			return cmp1 < 0
		}
		cmp2 := contracts[i].SlainAt.Compare(contracts[j].SlainAt)
		if cmp2 != 0 {
			return cmp2 < 0
		}
		return contracts[i].ContractID < contracts[j].ContractID
	})

	var maxDuration time.Duration
	for _, contract := range contracts {
		if contract.SlainAt.Sub(contract.TakenAt) > maxDuration {
			maxDuration = contract.SlainAt.Sub(contract.TakenAt)
		}
	}

	for _, contract := range contracts {
		if contract.SlainAt.Sub(contract.TakenAt) == maxDuration {
			difference := contract.SlainAt.Sub(contract.TakenAt)
			x := int(difference.Hours())
			y := int(difference.Minutes()) % 60

			fmt.Printf("Это была легендарная битва! Контракт %d был выполнен %s. Бой с монстром %q длился целых %d часов и %d минут!\n", contract.ContractID, contract.SlainAt.Format(layout), contract.Monster, x, y)
		} else {
			fmt.Printf("Контракт %d был выполнен %s.\n", contract.ContractID, contract.SlainAt.Format(layout))
		}
	}
}

func main() {
	contracts := []MonsterContract{
		{ContractID: 101, Monster: "Griffin", TakenAt: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC), SlainAt: time.Date(2025, time.May, 20, 14, 0, 0, 0, time.UTC)},
		{ContractID: 102, Monster: "Hydra", TakenAt: time.Date(2025, time.June, 15, 9, 30, 0, 0, time.UTC), SlainAt: time.Date(2025, time.June, 15, 11, 0, 0, 0, time.UTC)},
		{ContractID: 103, Monster: "Basilisk", TakenAt: time.Date(2025, time.May, 20, 10, 0, 0, 0, time.UTC), SlainAt: time.Date(2025, time.May, 20, 12, 30, 0, 0, time.UTC)},
		{ContractID: 104, Monster: "Wyvern", TakenAt: time.Date(2025, time.April, 10, 14, 0, 0, 0, time.UTC), SlainAt: time.Date(2025, time.April, 10, 18, 0, 0, 0, time.UTC)},
	}
	PrintChronicles(contracts)
}
