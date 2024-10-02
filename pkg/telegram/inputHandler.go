package telegram

//
//func ConstructAnswerMessage(answer *entity.Answer, bot *tgbotapi.BotAPI, db *gorm.DB) entity.Sendable {
//
//	messageConstructor := entity.MessageConstructor{}.New(&answer.NextMessage.Text, &answer.NextMessage.Type, answer.NextMessage.Keyboard, db)
//
//	var output = tgObjects.New(messageConstructor, bot)
//	return output
//}
//
//func SetChatMember(db *gorm.DB, chatMember tgbotapi.ChatMemberUpdated, bot models.Bot) {
//	if chatMember.OldChatMember.Status == "left" && chatMember.NewChatMember.Status == "administrator" {
//		botStandChatAdmin(db, chatMember, bot)
//	}
//	if chatMember.OldChatMember.Status == "administrator" && chatMember.NewChatMember.Status == "left" {
//		botLeftChannel(db, chatMember, bot)
//	}
//}
//
//func botStandChatAdmin(db *gorm.DB, chatMember tgbotapi.ChatMemberUpdated, bot models.Bot) {
//	var channel models.Channel
//	if err := db.Where("channel_tg_id = ?", chatMember.Chat.ID).Where("bot_id = ?", bot.Id).Find(&channel).Error; err != nil {
//		location, _ := time.LoadLocation("Europe/Moscow")
//		dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")
//		channel := models.Channel{
//			ChannelTgId: chatMember.Chat.ID,
//			Title:       chatMember.Chat.Title,
//			Username:    chatMember.Chat.UserName,
//			BotId:       bot.Id,
//			CreatedAt:   dateTime,
//			UpdatedAt:   dateTime,
//		}
//		db.Create(&channel)
//	}
//}
//
//func botLeftChannel(db *gorm.DB, chatMember tgbotapi.ChatMemberUpdated, bot models.Bot) {
//	var channel models.Channel
//	if err := db.Where("channel_tg_id = ?", chatMember.Chat.ID).Where("bot_id = ?", bot.Id).Find(&channel).Error; err == nil {
//		db.Delete(&channel)
//		fmt.Println("delete")
//	}
//}
