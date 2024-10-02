package entity

import (
	"github.com/Vlad06013/apiGin/models"
	"github.com/Vlad06013/apiGin/servises/repository"
	"github.com/jinzhu/gorm"
)

type TgUser struct {
	models.TgUser
}

type TgUserMessageHistory struct {
	models.TgUserMessageHistory
}

func InitUser(db *gorm.DB, tgID int64, name string) TgUser {
	var user = models.TgUser{
		TgUserId: tgID,
		Name:     name,
	}
	if err := db.Where("tg_user_id = ?", user.TgUserId).Preload("History").Find(&user).Error; err != nil {
		user = repository.CreateUser(db, &user)
	}
	userEntity := TgUser{user}
	return userEntity
}

func (u TgUser) GetBotHistory(db *gorm.DB, bot *Bot) TgUserMessageHistory {
	var history = models.TgUserMessageHistory{
		BotId:    bot.Id,
		TgUserId: u.Id,
	}
	historyExist, err := repository.GetMessageHistory(db, history.BotId, history.TgUserId)
	if err != nil {
		history = repository.CreateMessageHistory(db, &history)
	} else {
		history = historyExist
	}

	var historyEntity = TgUserMessageHistory{history}
	return historyEntity
}

func (u TgUser) GenerateAnswer(db *gorm.DB, bot *Bot) Answer {
	var lastMessage, nextMessage *models.Message
	history := u.GetBotHistory(db, bot)
	answer := Answer{
		User:         u,
		BotHistoryId: history.Id,
		ChatId:       u.TgUserId,
	}

	if history.LastMessageId != 0 {
		lastMessage, _ = repository.GetMessageById(db, history.LastMessageId)
		if lastMessage != nil {
			if lastMessage.NextMessageId != 0 {
				nextMessage, _ = repository.GetMessageById(db, lastMessage.NextMessageId)
			}
		}
	} else {
		firstMessage, err := repository.FirstMessage(db)
		if err == nil {
			nextMessage = firstMessage
		}
	}
	if lastMessage != nil {
		answer.LastMessage = Message{*lastMessage}
	}
	if nextMessage != nil {
		answer.NextMessage = Message{*nextMessage}
	}
	return answer
}

func (u TgUser) GenerateAnswerByCallbackData(db *gorm.DB, bot *Bot, pressedButton any) Answer {
	var lastMessage, nextMessage *models.Message
	history := u.GetBotHistory(db, bot)
	answer := Answer{
		User:         u,
		BotHistoryId: history.Id,
		ChatId:       u.TgUserId,
	}

	if history.LastMessageId != 0 {
		lastMessage = &history.LastMessage
	} else {
		firstMessage, err := repository.FirstMessage(db)
		if err == nil {
			nextMessage = firstMessage
		}
	}
	if lastMessage != nil {
		answer.LastMessage = Message{*lastMessage}
	}
	if nextMessage != nil {
		answer.NextMessage = Message{*nextMessage}
	} else {
		messagable := repository.GetMessagable(db, answer.LastMessage.Id, pressedButton)
		answer.NextMessage = Message{*messagable.ToMessage}
	}
	//fmt.Println(answer.LastMessage)
	//fmt.Println(answer.NextMessage)
	return answer
}

func (u TgUser) SaveLastMessage(db *gorm.DB, answer *Answer, LastTGMessageId int) TgUserMessageHistory {
	var history = models.TgUserMessageHistory{
		LastMessageId:   answer.NextMessage.Id,
		LastTGMessageId: LastTGMessageId,
	}
	history = repository.UpdateMessageHistory(db, answer.BotHistoryId, &history)
	historyEntity := TgUserMessageHistory{history}

	return historyEntity
}
