package telegram

import (
	"github.com/Vlad06013/apiGin/entity"
	"github.com/Vlad06013/apiGin/servises/constructor"
	"github.com/Vlad06013/apiGin/servises/repository"
)

func SendAnswer(toSend *entity.ToSend) {
	result := toSend.Output.SendMessage(toSend.Answer.ChatId)
	if result != nil {

		var history entity.TgUserMessageHistory
		if result.MessageID != 0 {
			history = toSend.Answer.User.SaveLastMessage(toSend.DB, &toSend.Answer, toSend.CallBackParsed, result.MessageID)
			toSend.Answer.User.BotHistory = &history
		}
		lastMessage, _ := repository.GetMessageById(toSend.DB, history.LastMessageId)
		var lastMessageEntity = entity.Message{Message: *lastMessage}
		if lastMessageEntity.CanSendNext() == true {
			nextAnswer, _ := toSend.Answer.User.GenerateAnswer(toSend.DB, toSend.Bot.Bot, nil)
			toSend.Answer = nextAnswer
			sendNextAnswer(toSend)
		}
	}

}

func sendNextAnswer(ToSend *entity.ToSend) {

	constructorParams := entity.ConstructorParams{
		Answer:  ToSend.Answer,
		BotApi:  ToSend.Bot.Api,
		DB:      ToSend.DB,
		Message: ToSend.Answer.NextMessage,
	}

	messageConstruct := constructor.ConstructAnswerMessage(&constructorParams)
	output := entity.NewOutput(&messageConstruct, &ToSend.Bot.Api)
	ToSend.Constructable = messageConstruct
	ToSend.Output = output
	SendAnswer(ToSend)
}
