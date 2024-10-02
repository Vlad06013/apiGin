package models

type Bot struct {
	Id    uint   `json:"id" gorm:"AUTO_INCREMENT;primary_key;column:id"`
	Token string `json:"token" gorm:"column:token"`
	Name  string `json:"name" gorm:"column:name"`
}
