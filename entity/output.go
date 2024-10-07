package entity

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Sendable interface {
	SendMessage
}

type SendMessage interface {
	SendMessage(chatId int64) *tgbotapi.Message
	DeleteMessage(chatId int64, messageID int) tgbotapi.Message
}
type Output struct {
	Constructable
	Bot tgbotapi.BotAPI
}

func (o *Output) sendTextMessage(chatId int64) *tgbotapi.Message {

	msg := tgbotapi.NewMessage(chatId, *o.Constructable.TextMessage())
	msg.ParseMode = "HTML"
	buttons := o.Constructable.ButtonsMessage()
	if len(buttons) != 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons)
	}

	res, err := o.Bot.Send(msg)
	if err != nil {
		fmt.Println("sendError", err)
	}

	return &res
}
func (o *Output) sendAlert(cbId *string, text string) tgbotapi.APIResponse {

	msg := tgbotapi.NewCallbackWithAlert(*cbId, text)

	res, err := o.Bot.Request(msg)

	if err != nil {
		fmt.Println("sendTGError", err)
	}

	return *res
}

//func (o *Output) sendAnimation(chatId int64) tgbotapi.Message {
//	file := tgbotapi.FileID("CgACAgIAAyEGAASPQOlFAAMHZvP7_TX74Su9L2n4ObpFez8brP0AAjFUAALjK6BLX0m8jZWiWzM2BA")
//	msg := tgbotapi.NewAnimation(chatId, file)
//	msg.Caption = o.Text
//	msg.ParseMode = "HTML"
//
//	var buttons []tgbotapi.InlineKeyboardButton
//	if len(o.Keyboard.Buttons) != 0 {
//		for _, b := range o.Keyboard.Buttons {
//			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(b.Text, b.CallbackData))
//		}
//		btns := tgbotapi.NewInlineKeyboardRow(buttons...)
//		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(btns)
//	}
//
//	res, _ := o.Bot.Send(msg)
//	return res
//}

//	func (o *Output) sendPhoto(chatId int64) tgbotapi.Message {
//		file := tgbotapi.FileID("CgACAgIAAyEGAASPQOlFAAMHZvP7_TX74Su9L2n4ObpFez8brP0AAjFUAALjK6BLX0m8jZWiWzM2BA")
//		msg := tgbotapi.NewPhoto(chatId, file)
//		msg.Caption = o.Text
//		msg.ParseMode = "HTML"
//
//		var buttons []tgbotapi.InlineKeyboardButton
//		if len(o.Keyboard.Buttons) != 0 {
//			for _, b := range o.Keyboard.Buttons {
//				buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(b.Text, b.CallbackData))
//			}
//			btns := tgbotapi.NewInlineKeyboardRow(buttons...)
//			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(btns)
//		}
//
//		res, _ := o.Bot.Send(msg)
//		return res
//	}
func (o *Output) DeleteMessage(chatId int64, messageID int) tgbotapi.Message {
	msg := tgbotapi.NewDeleteMessage(chatId, messageID)
	res, _ := o.Bot.Send(msg)
	return res
}
func (o *Output) SendMessage(chatId int64) *tgbotapi.Message {
	var sent *tgbotapi.Message
	if o.Constructable.TypeMessage() == "message" {
		sent = o.sendTextMessage(chatId)
	}
	//if o.Type == "alert" {
	//	o.sendAlert(o.MessageConstructor.CallBackID, o.MessageConstructor.Text)
	//	sent = nil
	//}

	//res := o.sendAnimation(chatId)
	return sent

}
func NewOutput(c *Constructable, bot *tgbotapi.BotAPI) Sendable {
	var output Sendable = &Output{*c, *bot}
	return output
}
