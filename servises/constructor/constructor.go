package constructor

import (
	"github.com/Vlad06013/apiGin/entity"
)

func ConstructAnswerMessage(constructorParams *entity.ConstructorParams) entity.MessageConstructor {

	if constructorParams.Answer.NextMessage != nil {
		return *entity.NewMessageConstructor(constructorParams)
	} else {
		if constructorParams.CallBackParsed != nil {
			return constructByParsed(constructorParams)
		}
	}

	return entity.MessageConstructor{}
}

func constructByParsed(constructorParams *entity.ConstructorParams) entity.MessageConstructor {
	var message entity.MessageConstructor

	if constructorParams.CallBackParsed.Pointer == "alert" {
		message = entity.MessageConstructor{
			Text:       "test",
			Type:       constructorParams.CallBackParsed.Pointer,
			Keyboard:   nil,
			Buttons:    nil,
			CallBackID: &constructorParams.CallBackParsed.Id,
		}
	}
	return message
}
