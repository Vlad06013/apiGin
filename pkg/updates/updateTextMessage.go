package updates

import (
	"github.com/Vlad06013/apiGin/entity"
	"github.com/Vlad06013/apiGin/pkg/telegram"
	"github.com/Vlad06013/apiGin/servises/constructor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func TextMessageHandler(db *gorm.DB, bot *entity.BotApi, message *tgbotapi.Message) {

	user := entity.InitUser(db, message.From.ID, message.From.UserName, bot.Bot)
	answer, _ := user.GenerateAnswer(db, bot.Bot, nil)
	constructorParams := entity.ConstructorParams{
		Answer:  answer,
		BotApi:  bot.Api,
		DB:      db,
		Message: &answer.NextMessage,
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
	telegram.SendAnswer(&toSend)

}
