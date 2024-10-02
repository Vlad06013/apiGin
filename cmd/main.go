package main

import (
	"fmt"
	"github.com/Vlad06013/apiGin/database"
	"github.com/Vlad06013/apiGin/entity"
	"github.com/Vlad06013/apiGin/pkg/updates"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
)

func main() {

	db := database.ConnectDB()
	var bots []entity.Bot

	var counts int
	db.Find(&bots).Count(&counts)
	var botApi = make([]*entity.BotApi, int(counts))
	if counts == 0 {
		panic("Ботов нет")
	}

	for i, bot := range bots {
		api, _ := tgbotapi.NewBotAPI(bot.Token)
		botApi[i] = &entity.BotApi{
			Api: *api,
			Bot: &bot,
		}
	}
	var wg sync.WaitGroup
	wg.Add(counts)

	work := func(bot *entity.BotApi) {
		defer wg.Done()
		fmt.Printf("Мониторинг бота %d старт \n", bot.Bot.Name)
		updates.CheckUpdates(bot, db)
	}
	for _, bot := range botApi {
		go work(bot)
	}
	wg.Wait()
}

//https://t.me/+t4gzaaZ8KDc0ODli
