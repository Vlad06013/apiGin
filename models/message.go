package models

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	Id               int    `json:"id" gorm:"AUTO_INCREMENT;primary_key;column:id"`
	Name             string `json:"name" gorm:"column:name"`
	Type             string `json:"type" gorm:"column:type"`
	Text             string `json:"text" gorm:"column:text"`
	FirstMessage     bool   `json:"first_message" gorm:"column:first_message"`
	WaitInput        string `json:"wait_input" gorm:"column:wait_input"`
	NeedConfirmation bool   `json:"need_confirmation" gorm:"column:need_confirmation"`
	BotId            uint   `json:"bot_id" gorm:"column:bot_id;foreign_key:bot_id"`
	NextMessageId    uint   `json:"next_message_id" gorm:"column:next_message_id;foreign_key:next_message_id"`

	NextMessage *Message `json:"next_message" gorm:"foreignKey:NextMessageId;default:null"`
	Keyboard    Keyboard
}

func FirstMessage(db *gorm.DB) Message {
	var firstMessage = Message{}
	db.Where("first_message = ?", true).Preload("Keyboard").Preload("Keyboard.Buttons").Find(&firstMessage)
	return firstMessage
}
