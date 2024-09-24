package models

type Buttons struct {
	Id           uint   `json:"id" gorm:"AUTO_INCREMENT;primary_key;column:id"`
	KeyboardId   uint   `json:"keyboard_id" gorm:"column:keyboard_id"`
	Text         string `json:"text" gorm:"column:text"`
	CallbackData string `json:"callback_data" gorm:"column:callback_data"`
}
