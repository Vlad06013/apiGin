package models

type Channel struct {
	Id          int    `json:"id" gorm:"AUTO_INCREMENT;primary_key;column:id"`
	Type        string `json:"type" gorm:"column:type"`
	ChannelTgId int64  `json:"channel_tg_id" gorm:"column:channel_tg_id"`
	Title       string `json:"title" gorm:"column:title"`
	Username    string `json:"username" gorm:"column:username"`
	BotId       uint   `json:"bot_id" gorm:"column:bot_id"`
	CreatedAt   any    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   any    `json:"updated_at" gorm:"column:updated_at"`
	Bot         *Bot   `json:"bot" gorm:"foreignKey:BotId;default:null"`
}
