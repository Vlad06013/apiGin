package models

type Keyboard struct {
	Id              uint   `json:"id" gorm:"AUTO_INCREMENT;primary_key;column:id"`
	Name            string `json:"name" gorm:"column:name"`
	MessageID       uint   `json:"message_id" gorm:"column:message_id"`
	ResizeKeyboard  bool   `json:"resize_keyboard" gorm:"column:resize_keyboard"`
	OneTimeKeyboard bool   `json:"one_time_keyboard" gorm:"column:one_time_keyboard"`
	Buttons         []Buttons
}
