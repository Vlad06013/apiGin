package constructor

import (
	"github.com/Vlad06013/apiGin/entity"
)

func ConstructAnswerMessage(constructorParams *entity.ConstructorParams) entity.MessageConstructor {
	messageConstructor := entity.NewMessageConstructor(constructorParams)
	return *messageConstructor
}
