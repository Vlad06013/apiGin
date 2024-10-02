package updates

import (
	"github.com/Vlad06013/apiGin/entity"
	"github.com/Vlad06013/apiGin/pkg/telegram"
	"github.com/Vlad06013/apiGin/servises/constructor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

//func ReadCallbackQuery(db *gorm.DB, bot *BotApi, callback *tgbotapi.CallbackQuery) models.Answer {
//	user := models.SetUser(db, callback.From.ID, callback.From.UserName)
//	answer := user.GenerateAnswerByCallbackData(db, bot.Bot, callback.Data)
//	return answer
//}

func CallbackQueryHandler(db *gorm.DB, bot *entity.BotApi, callbackQuery *tgbotapi.CallbackQuery) {

	user := entity.InitUser(db, callbackQuery.From.ID, callbackQuery.From.UserName)
	answer := user.GenerateAnswerByCallbackData(db, bot.Bot, callbackQuery.Data)
	messageConstruct := constructor.ConstructAnswerMessage(&answer, &bot.Api, db)
	output := entity.NewOutput(&messageConstruct, &bot.Api)
	toSend := entity.ToSend{
		answer,
		messageConstruct,
		output,
		db,
		bot,
	}
	telegram.SendAnswer(&toSend)
	//result := telegram.SendAnswer(output, answer, bot, db)

}
