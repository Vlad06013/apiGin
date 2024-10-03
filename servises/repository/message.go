package repository

import (
	"github.com/Vlad06013/apiGin/models"
	"github.com/jinzhu/gorm"
)

func GetMessageById(db *gorm.DB, messageId uint) (*models.Message, error) {

	var message models.Message
	err := db.Preload("Keyboard").Preload("Keyboard.Buttons").First(&message, messageId).Error
	return &message, err

}

func FirstMessage(db *gorm.DB) (*models.Message, error) {
	var firstMessage models.Message
	err := db.Where("first_message = ?", true).Preload("Keyboard").Preload("Keyboard.Buttons").Find(&firstMessage).Error
	return &firstMessage, err
}

func GetMessagable(db *gorm.DB, fromMessageId uint, callbackData any) models.TgMessagable {
	messagable := models.TgMessagable{}
	db.Where("from_message_id = ?", fromMessageId).
		Where("callback_data = ?", callbackData).
		Preload("ToMessage").
		Preload("ToMessage.Keyboard").
		Preload("ToMessage.Keyboard.Buttons").
		First(&messagable)

	return messagable
}
