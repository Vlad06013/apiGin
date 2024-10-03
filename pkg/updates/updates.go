package updates

import (
	"github.com/Vlad06013/apiGin/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func CheckUpdates(bot *entity.BotApi, db *gorm.DB) {
	botApi := bot.Api
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := botApi.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil {
			TextMessageHandler(db, bot, update.Message)
		}
		if update.CallbackQuery != nil {
			CallbackQueryHandler(db, bot, update.CallbackQuery)

			//	answer := ReadCallbackQuery(db, bot, update.CallbackQuery)
			//	telegram.Start(&answer, &botApi, db)
			//	//	go deleteLastMessage(output, answer)
			//	//telegram.SendAnswer(output, answer, bot)
		}
		//if update.MyChatMember != nil {
		//	ReadMyChatMember(db,)
		//	//	telegram.SetUser(db, update.MyChatMember.From.ID, update.MyChatMember.From.UserName)
		//	//	telegram.SetChatMember(db, *update.MyChatMember, *bot.Bot)
		//}
	}
}

//func SendAnswer(output Output.Sendable, answer models.Answer, bot *BotApi) tgbotapi.Message {
//	res := output.SendMessage(answer.ChatId)
//	if res.MessageID != 0 {
//		answer.User = answer.User.SaveLastMessage(db, &answer, res.MessageID)
//	}
//
//	fmt.Println(answer.CanSendNextMessage)
//	if answer.CanSendNextMessage == true {
//
//		nextAnswer := answer.User.GenerateAnswer(db, bot.Bot)
//		sendNextAnswer(&nextAnswer, bot)
//	}
//	return res
//
//}
//
//func sendNextAnswer(nextAnswer *models.Answer, bot *BotApi) {
//	output := telegram.Start(nextAnswer, &bot.Api, db)
//	SendAnswer(output, *nextAnswer, bot)
//}
