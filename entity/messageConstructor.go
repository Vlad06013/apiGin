package entity

import (
	"database/sql"
	"github.com/Vlad06013/apiGin/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

type MessageConstructor struct {
	Text     string
	Type     string
	Keyboard models.Keyboard
	Buttons  []tgbotapi.InlineKeyboardButton
}

func GenerateButtons(keyboard models.Keyboard, db *gorm.DB, callBackQuery *string) []tgbotapi.InlineKeyboardButton {
	var buttons, convertedButtons []tgbotapi.InlineKeyboardButton

	if len(keyboard.Buttons) != 0 {
		for _, b := range keyboard.Buttons {
			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(b.Text, b.CallbackData))
		}
		convertedButtons = tgbotapi.NewInlineKeyboardRow(buttons...)
	} else {

		if keyboard.TableName != "" {

			buttons = generateButtonsFromTable(&keyboard, db, callBackQuery)
			convertedButtons = tgbotapi.NewInlineKeyboardRow(buttons...)
		}
	}
	return convertedButtons
}
func generateButtonsFromTable(keyboard *models.Keyboard, db *gorm.DB, callBackQuery *string) []tgbotapi.InlineKeyboardButton {

	var buttons []tgbotapi.InlineKeyboardButton
	var rows *sql.Rows
	if callBackQuery != nil && keyboard.InputFilterField != "" {
		rows, _ = db.Table(keyboard.TableName).
			Where(keyboard.InputFilterField+"=?", callBackQuery).
			Select([]string{keyboard.KeyToButtonText, keyboard.KeyToButtonCallbackData}).
			Rows()
	} else {
		rows, _ = db.Table(keyboard.TableName).
			Select([]string{keyboard.KeyToButtonText, keyboard.KeyToButtonCallbackData}).
			Rows()
	}

	var buttonText string
	var callbackData string
	for rows.Next() {
		rows.Scan(&buttonText, &callbackData)
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(buttonText, callbackData))
	}
	return buttons
}

func NewMessageConstructor(constructorParams *ConstructorParams) *MessageConstructor {

	keyboard := constructorParams.Answer.NextMessage.Keyboard
	db := constructorParams.DB
	text := constructorParams.Answer.NextMessage.Text
	messageType := constructorParams.Answer.NextMessage.Message.Type
	callBackQuery := constructorParams.CallBackQuery
	buttons := GenerateButtons(keyboard, db, callBackQuery)
	var message = &MessageConstructor{text, messageType, keyboard, buttons}
	return message
}
