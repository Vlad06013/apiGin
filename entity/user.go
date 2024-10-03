package entity

import (
	"github.com/Vlad06013/apiGin/models"
	"github.com/Vlad06013/apiGin/servises/repository"
	"github.com/jinzhu/gorm"
)

type TgUser struct {
	models.TgUser
	BotHistory *TgUserMessageHistory
}

type TgUserMessageHistory struct {
	models.TgUserMessageHistory
	NextMessage *models.Message
}

func InitUser(db *gorm.DB, tgID int64, name string, bot *Bot) TgUser {
	var user = models.TgUser{
		TgUserId: tgID,
		Name:     name,
	}
	if err := db.Where("tg_user_id = ?", user.TgUserId).Find(&user).Error; err != nil {
		user = repository.CreateUser(db, &user)
	}
	userEntity := TgUser{user, nil}
	history := userEntity.GetBotHistory(db, bot)
	userEntity.BotHistory = &history
	return userEntity
}

func (u TgUser) GetBotHistory(db *gorm.DB, bot *Bot) TgUserMessageHistory {
	var history = models.TgUserMessageHistory{
		BotId:    bot.Id,
		TgUserId: u.Id,
	}
	var nextMessage models.Message
	historyExist, err := repository.GetMessageHistory(db, history.BotId, history.TgUserId)
	if err != nil {
		history = repository.CreateMessageHistory(db, &history)
	} else {
		history = historyExist
	}
	db.First(&nextMessage, history.LastMessage.NextMessageId)
	var historyEntity = TgUserMessageHistory{history, &nextMessage}
	return historyEntity
}

func (u TgUser) GenerateAnswer(db *gorm.DB, bot *Bot, pressedButton *string) (Answer, CallbackParsed) {

	answerGenerator := AnswerGenerator{
		User:         u,
		DB:           db,
		Bot:          *bot,
		History:      *u.BotHistory,
		CallBackData: pressedButton,
	}
	answer, callbackParsed := answerGenerator.GenerateAnswer()
	return answer, callbackParsed
}

func (u TgUser) SaveLastMessage(db *gorm.DB, answer *Answer, LastTGMessageId int) TgUserMessageHistory {
	var history = models.TgUserMessageHistory{
		LastMessageId:   answer.NextMessage.Id,
		LastTGMessageId: LastTGMessageId,
	}
	history = repository.UpdateMessageHistory(db, answer.User.BotHistory.Id, &history)
	historyEntity := TgUserMessageHistory{history, nil}

	return historyEntity
}
