package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type TgUser struct {
	Id     uint   `json:"id" gorm:"primary_key;column:id"`
	UserId any    `json:"user_id" gorm:"column:user_id;default:null"`
	Name   string `json:"name" gorm:"column:name;default:noname"`
	Email  string `json:"email" gorm:"column:email;default:null"`
	//LastMessageId int    `json:"last_message_id" gorm:"column:last_message_id;default:null;OnDelete:SET NULL;"`
	LastMessageId   int    `json:"last_message_id" gorm:"foreignKey:id;default:null;OnDelete:SET NULL;"`
	LastTGMessageId int    `json:"last_tg_message_id" gorm:"column:last_tg_message_id;default:null;OnDelete:SET NULL;"`
	Phone           any    `json:"phone" gorm:"column:phone;default:null"`
	TgUserId        int64  `json:"tg_user_id" gorm:"column:tg_user_id;unique"`
	TgUserName      string `json:"tg_user_name" gorm:"column:tg_user_name"`
	CreatedAt       any    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       any    `json:"updated_at" gorm:"column:updated_at"`
	//LastMessage Message `json:"last_message" gorm:"foreignKey:LastMessageId;default:null"`
}

func (u TgUser) SaveLastMessage(db *gorm.DB, LastMessageId int, LastTGMessageId int) TgUser {
	u.LastMessageId = LastMessageId
	u.LastTGMessageId = LastTGMessageId
	db.Save(&u)
	return u
}

func (u TgUser) GenerateAnswer(db *gorm.DB) Answer {
	var lastMessage Message
	var nextMessage Message

	answer := Answer{
		User:   u,
		ChatId: u.TgUserId,
	}

	if u.LastMessageId != 0 {
		db.Where("id = ?", u.LastMessageId).Preload("Keyboard").Preload("Keyboard.Buttons").Find(&lastMessage)
		answer.LastMessage = &lastMessage

	} else {
		nextMessage = FirstMessage(db)
		answer.NextMessage = &nextMessage
	}

	if lastMessage.NextMessageId != 0 {
		db.Where("id = ?", lastMessage.NextMessageId).Preload("Keyboard").Preload("Keyboard.Buttons").Find(&nextMessage)
		answer.NextMessage = &nextMessage
	}

	return answer
}

func (u TgUser) GenerateAnswerByCallbackData(db *gorm.DB, pressedButton any) Answer {
	fmt.Println(pressedButton)
	var lastMessage Message
	var nextMessage Message

	answer := Answer{
		User:   u,
		ChatId: u.TgUserId,
	}

	if u.LastMessageId != 0 {
		db.Where("id = ?", u.LastMessageId).Preload("Keyboard").Preload("Keyboard.Buttons").Find(&lastMessage)
		answer.LastMessage = &lastMessage

	} else {
		nextMessage = FirstMessage(db)
		answer.NextMessage = &nextMessage
	}

	if pressedButton != 0 {
		db.Where("id = ?", pressedButton).Preload("Keyboard").Preload("Keyboard.Buttons").Find(&nextMessage)
		answer.NextMessage = &nextMessage
	}

	return answer
}
