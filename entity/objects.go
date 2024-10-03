package entity

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

type BotApi struct {
	Api tgbotapi.BotAPI
	Bot *Bot
}

type ToSend struct {
	Answer             Answer
	MessageConstructor MessageConstructor
	Output             Sendable
	DB                 *gorm.DB
	Bot                *BotApi
}

type ConstructorParams struct {
	Answer        Answer
	BotApi        tgbotapi.BotAPI
	DB            *gorm.DB
	CallBackQuery *string
}
