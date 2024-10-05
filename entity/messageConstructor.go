package entity

import (
	"github.com/Vlad06013/apiGin/models"
	"github.com/Vlad06013/apiGin/servises/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

type CallbackParsed struct {
	Id           string
	CallBackData *string
	Filter       *string
	Pointer      string
	PointerID    string
}
type MessageConstructor struct {
	Text       string
	Type       string
	Keyboard   *models.Keyboard
	Buttons    []tgbotapi.InlineKeyboardButton
	CallBackID *string
}

var message *Message
var db *gorm.DB

func GenerateButtons(keyboard models.Keyboard, callBackParsed *CallbackParsed) []tgbotapi.InlineKeyboardButton {
	var buttons, convertedButtons []tgbotapi.InlineKeyboardButton

	if len(keyboard.Buttons) != 0 {
		for _, b := range keyboard.Buttons {
			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(b.Text, "mess_"+b.CallbackData))
		}
		convertedButtons = tgbotapi.NewInlineKeyboardRow(buttons...)
	} else {

		if keyboard.TableName != "" {

			buttons = generateButtonsFromTable(&keyboard, callBackParsed)
			convertedButtons = tgbotapi.NewInlineKeyboardRow(buttons...)
		}
	}
	return convertedButtons
}

func generateButtonsFromTable(keyboard *models.Keyboard, callbackParsed *CallbackParsed) []tgbotapi.InlineKeyboardButton {

	var buttons []tgbotapi.InlineKeyboardButton
	detailField := "description"

	if callbackParsed != nil && callbackParsed.Filter != nil && keyboard.InputFilterField != "" {
		buttons = generateButtonsFromQueryWithFilter(keyboard, callbackParsed.Filter, detailField)
	} else {
		buttons = generateButtonsFromQuery(keyboard)
	}

	callBackData := checkBackBtn()
	if callBackData != "" {
		buttons = addBackBtn(buttons, callBackData)
	}
	return buttons
}

func getDataPrefix() string {
	var callbackDataQuery, prefix string
	if message.NextMessageId != 0 {
		prefix = "mess"
		messageId := strconv.FormatUint(uint64(message.NextMessageId), 10)
		callbackDataQuery = prefix + "_" + messageId + "/filter_"

	} else {
		prefix = "query"
		callbackDataQuery = prefix + "_"
	}
	return callbackDataQuery
}

func generateButtonsFromQuery(keyboard *models.Keyboard) []tgbotapi.InlineKeyboardButton {
	var buttonText, callbackDataValue, callbackData string
	var buttons []tgbotapi.InlineKeyboardButton

	rows, _ := db.Table(keyboard.TableName).
		Select([]string{keyboard.KeyToButtonText, keyboard.KeyToButtonCallbackData}).Rows()
	if rows != nil {
		callbackDataQuery := getDataPrefix()
		for rows.Next() {
			rows.Scan(&buttonText, &callbackDataValue)
			callbackData = callbackDataQuery + callbackDataValue
			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(buttonText, callbackData))
		}
	}
	return buttons
}

func generateButtonsFromQueryWithFilter(keyboard *models.Keyboard, filter *string, detailField string) []tgbotapi.InlineKeyboardButton {
	var buttonText, callbackDataValue, callbackData string
	var buttons []tgbotapi.InlineKeyboardButton

	rows, _ := db.Table(keyboard.TableName).
		Select([]string{keyboard.KeyToButtonText, keyboard.KeyToButtonCallbackData}).
		Where(keyboard.InputFilterField+"=?", filter).Where(detailField+"!=?", "").Rows()

	if rows != nil {
		//callbackDataQuery := getDataPrefix()
		for rows.Next() {
			rows.Scan(&buttonText, &callbackDataValue)
			//callbackData = callbackDataQuery + callbackDataValue
			callbackData = "alert_" + callbackDataValue
			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(buttonText, callbackData))
		}
	}
	return buttons
}

func checkBackBtn() string {
	//add user last filter in history and check btn in last msg
	if message != nil {
		value := strconv.FormatUint(uint64(message.Id), 10)
		lastMessage, err := repository.GetMessageWithFilter(db, "next_message_id", value)
		if err == nil {
			if len(lastMessage.Keyboard.Buttons) != 0 || lastMessage.Keyboard.TableName != "" {
				return strconv.FormatUint(uint64(lastMessage.Id), 10)
			}
		} else {
			messagable := repository.GetMessagableByNextMessage(db, message.Id)
			return strconv.FormatUint(uint64(messagable.FromMessageId), 10)
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

func NewMessageConstructor(constructorParams *ConstructorParams) *MessageConstructor {

	keyboard := constructorParams.Answer.NextMessage.Keyboard
	db = constructorParams.DB
	text := constructorParams.Answer.NextMessage.Text
	messageType := constructorParams.Answer.NextMessage.Message.Type
	callBackParsed := constructorParams.CallBackParsed
	if constructorParams.Message != nil {
		message = constructorParams.Message
	}

	buttons := GenerateButtons(keyboard, callBackParsed)

	var message = &MessageConstructor{text, messageType, &keyboard, buttons, nil}
	return message
}
