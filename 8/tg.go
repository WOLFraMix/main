package main // ТГ бот который повторяет твои сообщения

import (
	"context"
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// ссылка на библиотеку для работы с Telegram Bot API
// которую можно установить через "go get github.com/mymmrac/telego"

func main() {
<<<<<<< HEAD
	botToken := "token" // токен можно взять в @BotFather
=======
	botToken := "token"
>>>>>>> 0c2bfa38ac0e1355fd252a3c547ca29eed2a7e9f

	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} // создаем нового бота с указанным токеном и включаем логирование

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // создаем контекст для управления жизненным циклом бота

	updates, err := bot.UpdatesViaLongPolling(ctx, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} // начинаем получать обновления от Telegram с помощью метода long polling

	for update := range updates {
		if update.Message != nil {
			chatID := tu.ID(update.Message.Chat.ID)
			_, err := bot.CopyMessage(ctx, &telego.CopyMessageParams{
				ChatID:     chatID,
				FromChatID: chatID,
				MessageID:  update.Message.MessageID,
			})
			if err != nil {
				fmt.Println(err)
			} // если полученное обновление содержит сообщение
			// то мы используем метод CopyMessage для отправки этого сообщения обратно
		}
	}
}
