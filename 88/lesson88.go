package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(GetUserParam("https://superuser.ru/orders?sort=amount"))
	fmt.Println(GetUserParam("https://site.ru/page?user=pavel&id=123"))
	fmt.Println(GetUserParam("https://nothing.else/matters?drumsuser=larsulrich&birthyear=1963"))
	fmt.Println(GetUserParam("https://choose.your/destiny?user=first&user=second"))
	fmt.Println(GetUserParam("https://choose.your/destiny?user=&user=second"))
	fmt.Println(GetUserParam("https://site.of.pepe/query?text=wtfa&debug=true&user=CoCoShne"))
	fmt.Println(GetUserParam("http://be.happy/smile"))
}

func GetUserParam(str string) string {
	// если нет ? или user= то сразу отсеиваем записи
	if !strings.Contains(str, "?") || !strings.Contains(str, "user=") {
		return "not found"
	}

	parts := strings.Split(str, "?")
	if len(parts) < 2 || parts[1] == "" {
		return "not found"
	}

	lessParts := strings.Split(parts[1], "&")
	for _, v := range lessParts {
		userParts := strings.SplitN(v, "=", 2)
		if len(userParts) != 2 {
			return "not found"
		}
		key, result := userParts[0], userParts[1]
		if key == "user" && result != "" {
			return result // здесь важно сразу вернуть первый результат чтобы цикл не искал следующих совпадений
		}
	}
	return "not found" // если за цикл ничего не нашлось
}
