package updates

import (
	"github.com/Vlad06013/apiGin/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func CallbackQueryHandler(db *gorm.DB, bot *entity.BotApi, callbackQuery *tgbotapi.CallbackQuery) {

	//fmt.Println(callbackQuery.Data)

	//user := entity.InitUser(db, callbackQuery.From.ID, callbackQuery.From.UserName, bot.Bot)
	//answer, callbackParsed := user.GenerateAnswer(db, bot.Bot, callbackQuery)
	//constructorParams := entity.ConstructorParams{
	//	Answer:         answer,
	//	BotApi:         bot.Api,
	//	DB:             db,
	//	CallBackParsed: &callbackParsed,
	//	Message:        answer.NextMessage,
	//}
	//
	//messageConstruct := constructor.ConstructAnswerMessage(&constructorParams)
	//
	//output := entity.NewOutput(&messageConstruct, &bot.Api)
	//toSend := entity.ToSend{
	//	answer,
	//	messageConstruct,
	//	output,
	//	db,
	//	bot,
	//	&callbackParsed,
	//}
	//if messageConstruct.Type != "alert" {
	//	go output.DeleteMessage(answer.ChatId, answer.User.BotHistory.LastTGMessageId)
	//}
	//go telegram.SendAnswer(&toSend)
}
