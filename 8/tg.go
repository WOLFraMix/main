package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := "8734711794:AAG1RHwLVeSGRxVcPkVzk1nIWoqm0XUfHio"

	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	updates, err := bot.UpdatesViaLongPolling(ctx, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
			}
		}
	}
}
