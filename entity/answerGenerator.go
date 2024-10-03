package entity

import (
	"fmt"
	"github.com/Vlad06013/apiGin/models"
	"github.com/Vlad06013/apiGin/servises/repository"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

type AnswerGenerator struct {
	User         TgUser
	DB           *gorm.DB
	Bot          Bot
	History      TgUserMessageHistory
	CallBackData *string
}

var cbParsed CallbackParsed

func (a AnswerGenerator) GenerateAnswer() (Answer, CallbackParsed) {
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
	return answer, cbParsed

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

	fmt.Println("inputdata", *a.CallBackData)

	if a.History.LastMessageId != 0 {
		lastMessage, _ = repository.GetMessageById(a.DB, a.History.LastMessageId)
	}

	if strings.Contains(*a.CallBackData, "_") {
		nextMessage = a.parseCallback()
	}

	//if strings.Contains(*a.CallBackData, "query") {
	//	data := strings.Split(*a.CallBackData, "_")
	//
	//	fmt.Println(words[0])
	//	fmt.Println(words[1])
	//
	//}
	//if a.History.LastMessageId != 0 {
	//	lastMessage, _ = repository.GetMessageById(a.DB, a.History.LastMessageId)
	//	if lastMessage != nil && lastMessage.NextMessageId != 0 {
	//		m, _ := repository.GetMessageById(a.DB, lastMessage.NextMessageId)
	//		if m != nil {
	//			nextMessage = m
	//		}
	//	}
	//} else {
	//	firstMessage, _ := repository.FirstMessage(a.DB)
	//	if firstMessage != nil {
	//		nextMessage = firstMessage
	//	}
	//}
	//if nextMessage == nil && lastMessage != nil {
	//	messagable := repository.GetMessagable(a.DB, lastMessage.Id, a.CallBackData)
	//	nextMessage = messagable.ToMessage
	//}
	return lastMessage, nextMessage
}
func (a AnswerGenerator) parseCallback() *models.Message {

	data := strings.Split(*a.CallBackData, "/")
	params := strings.Split(data[0], "_")

	cbParsed.Pointer = params[0]
	cbParsed.PointerID = params[1]

	if len(data) > 1 {
		paramsFilter := strings.Split(data[1], "_")
		if paramsFilter[0] == "filter" {
			cbParsed.Filter = &paramsFilter[1]
		}
	}

	var nextMessage *models.Message
	if cbParsed.Pointer == "query" {
		nextMessage = parseQueryBtn(&a)
	}
	if cbParsed.Pointer == "mess" {
		nextMessage = parseMessBtn(a.DB, &cbParsed.PointerID)
	}

	return nextMessage
}

func parseQueryBtn(a *AnswerGenerator) *models.Message {

	messagable := repository.GetMessagable(a.DB, *a.CallBackData)
	nextMessage := messagable.ToMessage
	return nextMessage
}

func parseMessBtn(db *gorm.DB, pointerID *string) *models.Message {
	messageId, _ := strconv.Atoi(*pointerID)
	nextMessage, _ := repository.GetMessageById(db, uint(messageId))
	return nextMessage
}
