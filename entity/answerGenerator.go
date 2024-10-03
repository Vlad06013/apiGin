package entity

import (
	"github.com/Vlad06013/apiGin/models"
	"github.com/Vlad06013/apiGin/servises/repository"
	"github.com/jinzhu/gorm"
)

type AnswerGenerator struct {
	User         TgUser
	DB           *gorm.DB
	Bot          Bot
	History      TgUserMessageHistory
	CallBackData *string
}

func (a AnswerGenerator) GenerateAnswer() Answer {
	var lastMessage, nextMessage *models.Message

	answer := Answer{
		User:   a.User,
		ChatId: a.User.TgUserId,
	}
	if a.CallBackData != nil {
		lastMessage, nextMessage = a.answerOnCallBackMessage()
	} else {
		lastMessage, nextMessage = a.answerOnTextMessage()
	}

	if lastMessage != nil {
		answer.LastMessage = Message{Message: *lastMessage}
	}
	if nextMessage != nil {
		answer.NextMessage = Message{Message: *nextMessage}
	}

	return answer

}

func (a AnswerGenerator) answerOnTextMessage() (*models.Message, *models.Message) {
	var lastMessage, nextMessage *models.Message

	if a.History.LastMessageId != 0 {
		lastMessage, _ = repository.GetMessageById(a.DB, a.History.LastMessageId)
		if lastMessage != nil {
			nextMessage, _ = repository.GetMessageById(a.DB, lastMessage.NextMessageId)
		}
	} else {
		firstMessage, _ := repository.FirstMessage(a.DB)
		if firstMessage != nil {
			nextMessage = firstMessage
		}
	}

	return lastMessage, nextMessage
}
func (a AnswerGenerator) answerOnCallBackMessage() (*models.Message, *models.Message) {
	var lastMessage, nextMessage *models.Message

	if a.History.LastMessageId != 0 {
		lastMessage, _ = repository.GetMessageById(a.DB, a.History.LastMessageId)
		if lastMessage != nil && lastMessage.NextMessageId != 0 {
			m, _ := repository.GetMessageById(a.DB, lastMessage.NextMessageId)
			if m != nil {
				nextMessage = m
			}
		}
	} else {
		firstMessage, _ := repository.FirstMessage(a.DB)
		if firstMessage != nil {
			nextMessage = firstMessage
		}
	}
	if nextMessage == nil && lastMessage != nil {
		messagable := repository.GetMessagable(a.DB, lastMessage.Id, a.CallBackData)
		nextMessage = messagable.ToMessage
	}
	return lastMessage, nextMessage
}
func checkInputFilterByButton() {

}
