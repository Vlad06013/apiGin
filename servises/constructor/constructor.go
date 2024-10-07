package constructor

import (
	"github.com/Vlad06013/apiGin/entity"
)

func ConstructAnswerMessage(constructorParams *entity.ConstructorParams) entity.Constructable {
	var message entity.Constructable

	if constructorParams.Answer.NextMessage != nil {

		if len(constructorParams.Answer.NextMessage.Keyboard.Buttons) == 0 &&
			constructorParams.Answer.NextMessage.Keyboard.TableName != "" {
			return entity.NewMessageConstructor(constructorParams)
		} else {
			return entity.NewDefaultMessage(constructorParams)
		}
		//	return *entity.NewMessageConstructor(constructorParams)
		//} else {
		//	if constructorParams.CallBackParsed != nil {
		//		return constructByParsed(constructorParams)
		//	}
	}

	return message
}

//func constructByParsed(constructorParams *entity.ConstructorParams) entity.MessageConstructor {
//	var message entity.MessageConstructor
//
//	if constructorParams.CallBackParsed.Pointer == "alert" {
//		message = constructAlert(constructorParams)
//	}
//	return message
//}
//
//func constructAlert(constructorParams *entity.ConstructorParams) entity.MessageConstructor {
//	var message entity.MessageConstructor
//
//	message = entity.MessageConstructor{
//		Text:       "test",
//		Type:       constructorParams.CallBackParsed.Pointer,
//		Keyboard:   nil,
//		Buttons:    nil,
//		CallBackID: &constructorParams.CallBackParsed.Id,
//	}
//	return message
//
//}
