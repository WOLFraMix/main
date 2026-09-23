package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

func main() {
	s1 := []string{
		"Ольга -> кошки",
		"Евгения -> собаки",
		"Ольга -> собаки",
		"Нина -> кошки",
		"Евгения -> кошки",
	}

	PrintVotesResults(s1)
}

func PrintVotesResults(votes []string) {
	// userLastVote хранит последний выбранный вариант для каждого пользователя
	userLastVote := make(map[string]string)
	// userOrder сохраняет порядок появления новых пользователей для вывода
	userOrder := []string{}
	// voteCount подсчитывает итоговое количество голосов за каждый вариант
	voteCount := make(map[string]int)
	// repeatUsers считает количество "лишних" голосов от одного пользователя
	repeatUsers := make(map[string]int)

	for _, v := range votes {
		parts := strings.Split(v, " -> ")
		user, vote := parts[0], parts[1]

		_, ok := userLastVote[user]
		if !ok {
			// Если пользователь встречается впервые, фиксируем его появление
			userOrder = append(userOrder, user)
		} else {
			// Если голос уже был, увеличиваем счетчик повторов этого пользователя
			repeatUsers[user]++
		}
		// Перезаписываем значение — нам важен только самый последний выбор
		userLastVote[user] = vote
	}

	// Подсчёт финальных голосов
	for _, user := range userOrder {
		vote := userLastVote[user]
		voteCount[vote]++
	}

	if len(userOrder) == 0 {
		fmt.Println("Уникальные пользователи: -")
	} else {
		fmt.Printf("Уникальные пользователи: %s (всего %d)\n", strings.Join(userOrder, ", "), len(userOrder))
	}

	if len(repeatUsers) == 0 {
		fmt.Println("Повторные голоса: -")
	} else {
		repeatVotes := []string{}
		for _, user := range userOrder {
			count, ok := repeatUsers[user]
			if ok {
				repeatVotes = append(repeatVotes, fmt.Sprintf("%s (%d)", user, count+1))
			}
		}
		fmt.Printf("Повторные голоса: %s\n", strings.Join(repeatVotes, ", "))
	}

	if len(voteCount) == 0 {
		fmt.Println("Выиграло по голосам: -")
		return
	} else {
		maxVotes := 0
		for _, count := range voteCount {
			if count > maxVotes {
				maxVotes = count
			}
		}
		winners := []string{}
		for vote, count := range voteCount {
			if count == maxVotes {
				winners = append(winners, vote)
			}
		}
		if len(winners) == 1 {
			fmt.Printf("Выиграло по голосам: %s (%d)\n", winners[0], maxVotes)
		} else {
			cl := collate.New(language.Russian)
			cl.SortStrings(winners)
			fmt.Printf("Выиграло по голосам: %s (по %d)\n", strings.Join(winners, ", "), maxVotes)
		}
	}
}
