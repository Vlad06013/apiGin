package entity

import (
	"github.com/Vlad06013/apiGin/models"
	"github.com/Vlad06013/apiGin/servises/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

type MessageConstructor struct {
	Text       string
	Type       string
	Keyboard   *models.Keyboard
	Buttons    []tgbotapi.InlineKeyboardButton
	CallBackID *string
}

var nextMessage *Message
var db *gorm.DB

func (m *MessageConstructor) TextMessage() *string {
	return &m.Text
}
func (m *MessageConstructor) TypeMessage() string {
	return m.Type
}

func (m *MessageConstructor) ButtonsMessage() []tgbotapi.InlineKeyboardButton {
	return m.Buttons
}
func generateButtons(keyboard models.Keyboard) []tgbotapi.InlineKeyboardButton {
	var buttons, convertedButtons []tgbotapi.InlineKeyboardButton
	prefix := "mess"

	if len(keyboard.Buttons) != 0 {
		for _, b := range keyboard.Buttons {
			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(b.Text, prefix+"_"+b.CallbackData))
		}
		convertedButtons = tgbotapi.NewInlineKeyboardRow(buttons...)
	}
	callBackData := checkBackBtn()
	if callBackData != "" {
		convertedButtons = addBackBtn(convertedButtons, callBackData)
	}
	return convertedButtons
}

func checkBackBtn() string {
	if nextMessage != nil {
		value := strconv.FormatUint(uint64(nextMessage.Id), 10)
		lastMessage, err := repository.GetMessageWithFilter(db, "next_message_id", value)
		if err == nil && len(lastMessage.Keyboard.Buttons) != 0 {
			return strconv.FormatUint(uint64(lastMessage.Id), 10)
		}
	}
	return ""
}

func addBackBtn(buttons []tgbotapi.InlineKeyboardButton, callback string) []tgbotapi.InlineKeyboardButton {
	return addCustomBtn(buttons, "Назад", "mess_"+callback)
}

func addCustomBtn(buttons []tgbotapi.InlineKeyboardButton, text string, callback string) []tgbotapi.InlineKeyboardButton {
	return append(buttons, tgbotapi.NewInlineKeyboardButtonData(text, callback))
}

func NewDefaultMessage(constructorParams *ConstructorParams) Constructable {

	keyboard := constructorParams.Answer.NextMessage.Keyboard
	db = constructorParams.DB
	text := constructorParams.Answer.NextMessage.Text
	messageType := constructorParams.Answer.NextMessage.Message.Type
	if constructorParams.Message != nil {
		nextMessage = constructorParams.Message
	}

	buttons := generateButtons(keyboard)

	var defaultMessage Constructable = &MessageConstructor{
		Text:       text,
		Type:       messageType,
		Keyboard:   &keyboard,
		Buttons:    buttons,
		CallBackID: nil,
	}
	return defaultMessage
}
