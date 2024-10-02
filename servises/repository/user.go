package repository

import (
	"github.com/Vlad06013/apiGin/models"
	"github.com/jinzhu/gorm"
	"time"
)

func CreateUser(db *gorm.DB, u *models.TgUser) models.TgUser {

	location, _ := time.LoadLocation("Europe/Moscow")
	dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")

	user := models.TgUser{
		TgUserId:   u.TgUserId,
		TgUserName: u.Name,
		Name:       u.Name,
		CreatedAt:  dateTime,
		UpdatedAt:  dateTime,
	}
	db.Create(&user)
	return user
}
func GetMessageHistory(db *gorm.DB, botId uint, TgUserId uint) (models.TgUserMessageHistory, error) {

	var history models.TgUserMessageHistory
	err := db.Table("tg_user_message_histories").Preload("LastMessage").Where("bot_id = ?", &botId).Where("tg_user_id = ?", &TgUserId).First(&history).Error
	return history, err

}
func CreateMessageHistory(db *gorm.DB, h *models.TgUserMessageHistory) models.TgUserMessageHistory {

	history := models.TgUserMessageHistory{
		BotId:    h.BotId,
		TgUserId: h.TgUserId,
	}
	db.Create(&history)
	return history
}

func UpdateMessageHistory(db *gorm.DB, id uint, fields *models.TgUserMessageHistory) models.TgUserMessageHistory {
	history := models.TgUserMessageHistory{Id: id}

	db.First(&history)

	history.LastMessageId = fields.LastMessageId
	history.LastTGMessageId = fields.LastTGMessageId

	db.Save(&history)
	db.Preload("LastMessage").First(&history)

	return history
}
