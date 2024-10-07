package entity

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Constructable interface {
	TypeMessage() string
	TextMessage() *string
	ButtonsMessage() []tgbotapi.InlineKeyboardButton
}
