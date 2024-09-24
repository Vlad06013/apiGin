package Input

import (
	"github.com/Vlad06013/apiGin/models"
	"time"
)

type TextMessage struct {
	UpdateId int `json:"update_id"`
	Message  struct {
		MessageId int    `json:"message_id"`
		From      From   `json:"from"`
		Chat      Chat   `json:"chat"`
		Date      int    `json:"date"`
		Text      string `json:"text"`
	} `json:"message"`
}

func (m TextMessage) User() models.TgUser {
	var user models.TgUser
	var db = models.ConnectDB()
	//if err := db.Where("tg_user_id = ?", m.Message.From.ID).Preload("LastMessage").Preload("LastMessage.NextMessage").Find(&user).Error; err != nil {
	if err := db.Where("tg_user_id = ?", m.Message.From.ID).Find(&user).Error; err != nil {
		location, _ := time.LoadLocation("Europe/Moscow")
		dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")
		user := models.TgUser{
			TgUserId:   m.Message.From.ID,
			TgUserName: m.Message.From.UserName,
			Name:       m.Message.From.UserName,
			CreatedAt:  dateTime,
			UpdatedAt:  dateTime,
		}
		db.Create(&user)
		return user
	}
	return user
}

func (m TextMessage) ChatId() int {
	return m.Message.Chat.ID
}

func (m TextMessage) GenerateAnswer() models.Answer {
	var user = m.User()
	var lastMessage models.Message
	var nextMessage models.Message
	var db = models.ConnectDB()

	answer := models.Answer{
		User:   user,
		ChatId: m.ChatId(),
	}

	if user.LastMessageId != 0 {
		db.Where("id = ?", user.LastMessageId).Preload("Keyboard").Preload("Keyboard.Buttons").Find(&lastMessage)
		answer.LastMessage = &lastMessage

	} else {
		nextMessage = models.FirstMessage(db)
		answer.NextMessage = &nextMessage
	}

	if lastMessage.NextMessageId != 0 {
		db.Where("id = ?", lastMessage.NextMessageId).Preload("Keyboard").Preload("Keyboard.Buttons").Find(&nextMessage)
		answer.NextMessage = &nextMessage
	}

	return answer
}
