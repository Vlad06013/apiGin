package entity

import (
	"github.com/Vlad06013/apiGin/models"
)

//const modelNameSpace = "Valibool\\TelegramConstruct\\Models\\Message"

type Message struct {
	models.Message
}
type Messagable struct {
	models.TgMessagable
}

func (m Message) CanSendNext() bool {
	if len(m.Keyboard.Buttons) == 0 && m.NextMessageId != 0 && m.Keyboard.TableName == "" {
		return true
	}
	return false
}
