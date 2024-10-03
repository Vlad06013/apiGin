package updates

import (
	"github.com/Vlad06013/apiGin/entity"
	"github.com/Vlad06013/apiGin/pkg/telegram"
	"github.com/Vlad06013/apiGin/servises/constructor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func CallbackQueryHandler(db *gorm.DB, bot *entity.BotApi, callbackQuery *tgbotapi.CallbackQuery) {

	user := entity.InitUser(db, callbackQuery.From.ID, callbackQuery.From.UserName, bot.Bot)
	answer := user.GenerateAnswer(db, bot.Bot, &callbackQuery.Data)
	constructorParams := entity.ConstructorParams{
		Answer:        answer,
		BotApi:        bot.Api,
		DB:            db,
		CallBackQuery: &callbackQuery.Data,
	}

	messageConstruct := constructor.ConstructAnswerMessage(&constructorParams)
	output := entity.NewOutput(&messageConstruct, &bot.Api)
	toSend := entity.ToSend{
		answer,
		messageConstruct,
		output,
		db,
		bot,
	}

	go output.DeleteMessage(answer.ChatId, answer.User.BotHistory.LastTGMessageId)
	go telegram.SendAnswer(&toSend)
}
