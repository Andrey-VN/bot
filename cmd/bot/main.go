package main

import (
	"log"

	"github.com/Andrey-VN/bot/internal/app/commands"
	"github.com/Andrey-VN/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const token = "6230019625:AAGeXKRhPWRBNbGmkNf2hblyUm4YnN2H_Co"

func main() {
	// token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	// if err!=nil {
	// 	log.Panic(err)
	// }

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}

		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)
		case "list":
			commander.List(update.Message)
		default:
			commander.Default(update.Message)
		}
	}
}
