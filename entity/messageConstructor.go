package entity

import (
	"database/sql"
	"fmt"
	"github.com/Vlad06013/apiGin/models"
	"github.com/Vlad06013/apiGin/servises/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

type MessageConstructor struct {
	Text     string
	Type     string
	Keyboard models.Keyboard
	Buttons  []tgbotapi.InlineKeyboardButton
}

var message *Message
var db *gorm.DB

func GenerateButtons(keyboard models.Keyboard, callBackQuery *string) []tgbotapi.InlineKeyboardButton {
	var buttons, convertedButtons []tgbotapi.InlineKeyboardButton

	if len(keyboard.Buttons) != 0 {
		for _, b := range keyboard.Buttons {
			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(b.Text, b.CallbackData))
		}
		convertedButtons = tgbotapi.NewInlineKeyboardRow(buttons...)
	} else {

		if keyboard.TableName != "" {

			buttons = generateButtonsFromTable(&keyboard, callBackQuery)
			convertedButtons = tgbotapi.NewInlineKeyboardRow(buttons...)
		}
	}
	return convertedButtons
}
func generateButtonsFromTable(keyboard *models.Keyboard, callBackQuery *string) []tgbotapi.InlineKeyboardButton {

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
	callBackData := checkBackBtn()
	fmt.Println("backData", callBackData)
	if callBackData != "" {
		buttons = addBackBtn(buttons, callBackData)
	}
	return buttons
}

func checkBackBtn() string {
	if message != nil {
		value := strconv.FormatUint(uint64(message.Id), 10)
		lastMessage, err := repository.GetMessageWithFilter(db, "next_message_id", value)
		if err == nil {
			return strconv.FormatUint(uint64(lastMessage.Id), 10)
		} else {
			messagable := repository.GetMessagableByNextMessage(db, message.Id)
			return strconv.FormatUint(uint64(messagable.FromMessageId), 10)
		}
	}
	return ""
}
func addBackBtn(buttons []tgbotapi.InlineKeyboardButton, callback string) []tgbotapi.InlineKeyboardButton {
	return addCustomBtn(buttons, "Назад", callback)
}

func addCustomBtn(buttons []tgbotapi.InlineKeyboardButton, text string, callback string) []tgbotapi.InlineKeyboardButton {
	return append(buttons, tgbotapi.NewInlineKeyboardButtonData(text, callback))
}
func NewMessageConstructor(constructorParams *ConstructorParams) *MessageConstructor {

	keyboard := constructorParams.Answer.NextMessage.Keyboard
	db = constructorParams.DB
	text := constructorParams.Answer.NextMessage.Text
	messageType := constructorParams.Answer.NextMessage.Message.Type
	callBackQuery := constructorParams.CallBackQuery
	if constructorParams.Message != nil {
		message = constructorParams.Message
	}
	buttons := GenerateButtons(keyboard, callBackQuery)
	var message = &MessageConstructor{text, messageType, keyboard, buttons}
	return message
}
