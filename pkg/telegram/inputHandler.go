package telegram

import (
	"github.com/Vlad06013/apiGin/models"
	"github.com/Vlad06013/apiGin/models/tgObjects"
	"github.com/Vlad06013/apiGin/models/tgObjects/Output"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"time"
)

func Start(answer *models.Answer, bot *tgbotapi.BotAPI) Output.Sendable {
	messageConstructor := tgObjects.MessageConstructor{
		answer.NextMessage.Text,
		answer.NextMessage.Type,
		answer.NextMessage.Keyboard,
	}
	var output = Output.New(&messageConstructor, bot)
	return output
}
func SetUser(db *gorm.DB, tgID int64, name string) models.TgUser {
	var user models.TgUser
	if err := db.Where("tg_user_id = ?", tgID).Find(&user).Error; err != nil {
		location, _ := time.LoadLocation("Europe/Moscow")
		dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")
		user := models.TgUser{
			TgUserId:   tgID,
			TgUserName: name,
			Name:       name,
			CreatedAt:  dateTime,
			UpdatedAt:  dateTime,
		}
		db.Create(&user)
		return user
	}
	return user
}
