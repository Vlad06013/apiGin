package models

type Bot struct {
	Id    int    `json:"id" gorm:"AUTO_INCREMENT;primary_key;column:id"`
	Token string `json:"token"`
}
