package main

import (
	"github.com/Vlad06013/apiGin/models"
	"github.com/Vlad06013/apiGin/models/tgObjects/Output"
	"github.com/Vlad06013/apiGin/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var bot, err = tgbotapi.NewBotAPI("5492480377:AAF8-HRrtGtXIjaNhSaG9zt3hOnPk_MCuxg")

func main() {
	//bot, err := tgbotapi.NewBotAPI("6685188155:AAFfQxYZBwyC3mF-VfuKt6Nr9M-TtgIsa9s")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			user := telegram.SetUser(update.Message.From.ID, update.Message.From.UserName)
			answer := user.GenerateAnswer()
			output := telegram.Start(&answer, bot)
			SendAnswer(output, answer)

		}
		if update.CallbackQuery != nil {
			user := telegram.SetUser(update.CallbackQuery.From.ID, update.CallbackQuery.From.UserName)
			answer := user.GenerateAnswerByCallbackData(update.CallbackQuery.Data)
			output := telegram.Start(&answer, bot)
			SendAnswer(output, answer)
		}
	}
	//router := gin.Default()
	//
	//routes.InitRoutes(router)
	//models.ConnectDB()

	//router.Run(":8080")
}
func SendAnswer(output Output.Sendable, answer models.Answer) tgbotapi.Message {
	res := output.SendMessage(answer.ChatId)
	if res.MessageID != 0 {
		answer.User = answer.User.SaveLastMessage(answer.NextMessage.Id, res.MessageID)
	}
	nextAnswer := answer.User.GenerateAnswer()

	if nextAnswer.NextMessage != nil {
		sendNextAnswer(&nextAnswer)
	}
	return res

}

func sendNextAnswer(nextAnswer *models.Answer) {
	output := telegram.Start(nextAnswer, bot)
	SendAnswer(output, *nextAnswer)
}
