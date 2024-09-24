package Input

import (
	"github.com/Vlad06013/apiGin/models"
)

type Telegram interface {
	Chatable
	User
	MessageGenerator
}
type MessageGenerator interface {
	GenerateAnswer() models.Answer
}
type Chatable interface {
	ChatId() int
}

type User interface {
	User() models.TgUser
}

type From struct {
	ID        any    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	UserName  string `json:"username"`
}

type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	UserName  string `json:"username"`
}

type CallBackQuery struct {
	UpdateId      int `json:"update_id"`
	CallBackQuery struct {
		ID int `json:"id"`
	} `json:"callback_query"`
}
