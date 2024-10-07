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
			//CallbackQueryHandler(db, bot, update.CallbackQuery)
		}
		//if update.MyChatMember != nil {
		//	ReadMyChatMember(db,)
		//	//	telegram.SetUser(db, update.MyChatMember.From.ID, update.MyChatMember.From.UserName)
		//	//	telegram.SetChatMember(db, *update.MyChatMember, *bot.Bot)
		//}
	}
}
