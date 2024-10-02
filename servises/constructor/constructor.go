package constructor

import (
	"github.com/Vlad06013/apiGin/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func ConstructAnswerMessage(answer *entity.Answer, bot *tgbotapi.BotAPI, db *gorm.DB) entity.MessageConstructor {

	messageConstructor := entity.NewMessageConstructor(&answer.NextMessage.Text, &answer.NextMessage.Type, answer.NextMessage.Keyboard, db)
	return *messageConstructor
}
