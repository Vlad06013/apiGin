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

func GetMessageWithFilter(db *gorm.DB, field string, value any) (*models.Message, error) {
	var message models.Message
	//err := db.Where(field+" = ?", value).Preload("Keyboard").Preload("Keyboard.Buttons").First(&message).Error
	err := db.Where(field+" = ?", value).First(&message).Error
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

func GetMessagableByNextMessage(db *gorm.DB, toMessageId uint) models.TgMessagable {
	messagable := models.TgMessagable{}
	db.Where("to_message_id = ?", toMessageId).
		//Where("callback_data = ?", callbackData).
		Preload("ToMessage").
		Preload("ToMessage.Keyboard").
		Preload("ToMessage.Keyboard.Buttons").
		First(&messagable)

	return messagable
}
