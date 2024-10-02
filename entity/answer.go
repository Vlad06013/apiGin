package entity

type Answer struct {
	User         TgUser
	ChatId       int64
	LastMessage  Message
	NextMessage  Message
	BotHistoryId uint
}
