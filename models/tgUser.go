package models

type TgUser struct {
	Id         uint   `json:"id" gorm:"primary_key;column:id"`
	UserId     any    `json:"user_id" gorm:"column:user_id;default:null"`
	Name       string `json:"name" gorm:"column:name;default:noname"`
	Email      string `json:"email" gorm:"column:email;default:null"`
	Phone      any    `json:"phone" gorm:"column:phone;default:null"`
	TgUserId   int64  `json:"tg_user_id" gorm:"foreignKey:id;column:tg_user_id;unique"`
	TgUserName string `json:"tg_user_name" gorm:"column:tg_user_name"`
	CreatedAt  any    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  any    `json:"updated_at" gorm:"column:updated_at"`

	//History []TgUserMessageHistory
}

type TgUserMessageHistory struct {
	Id              uint    `json:"id" gorm:"primary_key;column:id"`
	TgUserId        uint    `json:"tg_user_id" gorm:"column:tg_user_id"`
	BotId           uint    `json:"bot_id" gorm:"column:bot_id"`
	LastMessageId   uint    `json:"last_message_id" gorm:"foreignKey:id;default:null;OnDelete:SET NULL;"`
	LastTGMessageId int     `json:"last_tg_message_id" gorm:"column:last_tg_message_id;default:null;OnDelete:SET NULL;"`
	LastQueryFilter string  `json:"last_query_filter" gorm:"column:last_query_filter;default:null;OnDelete:SET NULL;"`
	LastMessage     Message `json:"last_message" gorm:"foreignKey:LastMessageId;default:null"`
}

//func SetUser(db *gorm.DB, tgID int64, name string) TgUser {
//	var user TgUser
//	if err := db.Where("tg_user_id = ?", tgID).Preload("History").Find(&user).Error; err != nil {
//		location, _ := time.LoadLocation("Europe/Moscow")
//		dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")
//		user := TgUser{
//			TgUserId:   tgID,
//			TgUserName: name,
//			Name:       name,
//			CreatedAt:  dateTime,
//			UpdatedAt:  dateTime,
//		}
//		db.Create(&user)
//		return user
//	}
//	return user
//}

//func (u TgUser) SaveLastMessage(db *gorm.DB, answer *entity.Answer, LastTGMessageId int) TgUser {
//	var history = TgUserMessageHistory{Id: answer.BotHistoryId}
//	db.First(&history)
//	history.LastMessageId = answer.NextMessage.Id
//	//history.LastMessageId = 1
//	history.LastTGMessageId = LastTGMessageId
//	db.Save(&history)
//	return u
//}
//
//func (u TgUser) getBotHistory(db *gorm.DB, bot *Bot) *TgUserMessageHistory {
//	var history TgUserMessageHistory
//	if err := db.Table("tg_user_message_histories").Where("bot_id = ?", &bot.Id).Where("tg_user_id = ?", &u.Id).First(&history).Error; err != nil {
//		history := TgUserMessageHistory{
//			BotId:    bot.Id,
//			TgUserId: u.Id,
//		}
//		db.Create(&history)
//		return &history
//
//	}
//	return &history
//}
//
//func (u TgUser) GenerateAnswer(db *gorm.DB, bot *Bot) entity.Answer {
//	var nextMessage Message
//	var lastMessage Message
//	history := u.getBotHistory(db, bot)
//
//	answer := entity.Answer{
//		User:         u,
//		BotHistoryId: history.Id,
//		ChatId:       u.TgUserId,
//	}
//
//	if history.LastMessageId != 0 {
//		//db.Where("id = ?", u.LastMessageId).Preload("Keyboard").Preload("Keyboard.Buttons").Find(&lastMessage)
//		db.Preload("Keyboard").Preload("Keyboard.Buttons").First(&lastMessage, history.LastMessageId)
//		answer.LastMessage = &lastMessage
//		answer.CanSendNextMessage = canSendNext(&lastMessage)
//
//	} else {
//		nextMessage = FirstMessage(db)
//		//nextMessage.Attachments(db)
//		answer.NextMessage = &nextMessage
//		answer.CanSendNextMessage = canSendNext(&nextMessage)
//
//	}
//	//
//	if lastMessage.NextMessageId != 0 {
//		//db.Where("id = ?", lastMessage.NextMessageId).Preload("Keyboard").Preload("Keyboard.Buttons").Find(&nextMessage)
//		db.Preload("Keyboard").Preload("Keyboard.Buttons").First(&nextMessage, lastMessage.NextMessageId)
//
//		answer.NextMessage = &nextMessage
//		answer.CanSendNextMessage = canSendNext(&nextMessage)
//
//	}
//	return answer
//}
//
//func (u TgUser) GenerateAnswerByCallbackData(db *gorm.DB, bot *Bot, pressedButton any) entity.Answer {
//	var nextMessage Message
//	var lastMessage Message
//	history := u.getBotHistory(db, bot)
//
//	answer := entity.Answer{
//		User:         u,
//		BotHistoryId: history.Id,
//		ChatId:       u.TgUserId,
//	}
//
//	if history.LastMessageId != 0 {
//		db.Preload("Keyboard").Preload("Keyboard.Buttons").First(&lastMessage, history.LastMessageId)
//		answer.LastMessage = &lastMessage
//		answer.CanSendNextMessage = canSendNext(&lastMessage)
//
//	} else {
//		nextMessage = FirstMessage(db)
//		answer.NextMessage = &nextMessage
//		answer.CanSendNextMessage = canSendNext(&nextMessage)
//	}
//
//	if answer.NextMessage == nil {
//		messagable := TgMessagable{}
//		db.Where("from_message_id = ?", answer.LastMessage.Id).
//			Where("callback_data = ?", pressedButton).
//			Preload("ToMessage").
//			Preload("ToMessage.Keyboard").
//			Preload("ToMessage.Keyboard.Buttons").
//			First(&messagable, pressedButton)
//
//		answer.NextMessage = messagable.ToMessage
//	}
//	return answer
//}
//
//func canSendNext(message *Message) bool {
//	fmt.Println("buttons", message.Keyboard.Buttons, "NextMessageId", message.NextMessageId, "TableName", message.Keyboard.TableName)
//	//fmt.Println("NextMessageId",message.NextMessageId)
//	//fmt.Println("TableName",message.Keyboard.TableName)
//	if len(message.Keyboard.Buttons) == 0 && message.NextMessageId != 0 && message.Keyboard.TableName == "" {
//		return true
//	}
//	return false
//}
