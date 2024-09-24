package Output

import (
	"github.com/Vlad06013/apiGin/models/tgObjects"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Response map[string]interface{}

type Sendable interface {
	SendMessage
}

type SendMessage interface {
	SendMessage(chatId int64) tgbotapi.Message
}
type Output struct {
	tgObjects.MessageConstructor
	Bot tgbotapi.BotAPI
}

func (o *Output) sendTextMessage(chatId int64) tgbotapi.Message {
	msg := tgbotapi.NewMessage(chatId, o.Text)

	var buttons []tgbotapi.InlineKeyboardButton
	if len(o.Keyboard.Buttons) != 0 {
		for _, b := range o.Keyboard.Buttons {
			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(b.Text, b.CallbackData))
		}
		btns := tgbotapi.NewInlineKeyboardRow(buttons...)
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(btns)
	}

	res, _ := o.Bot.Send(msg)
	return res
}

func (o *Output) SendMessage(chatId int64) tgbotapi.Message {

	res := o.sendTextMessage(chatId)
	return res
	//body := url.Values{
	//	"chat_id":    {strconv.Itoa(chatId)},
	//	"text":       {o.MessageConstructor.Text},
	//	"parse_mode": {"HTML"},
	//}
	//resp, _ := http.PostForm(
	//	"https://api.telegram.org/bot5492480377:AAF8-HRrtGtXIjaNhSaG9zt3hOnPk_MCuxg/sendMessage",
	//	body,
	//)
	////if err != nil {
	////	fmt.Println(err)
	////	return
	////}
	//defer resp.Body.Close()
	//bodys, _ := ioutil.ReadAll(resp.Body)
	//
	//var response Response
	//err := json.Unmarshal([]byte(bodys), &response)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Printf("%s\n", response["result"])

	//c := json.NewDecoder(resp.Body)

	//fmt.Println(&c)

	//for true {
	//
	//	bs := make([]byte, 1014)
	//n, err := resp.Body.Read(bs)
	//fmt.Println(string(bs[:n]))
	//
	//	if n == 0 || err != nil {
	//		break
	//	}
	//}
}

func New(m *tgObjects.MessageConstructor, bot *tgbotapi.BotAPI) Sendable {
	var output Sendable = &Output{*m, *bot}
	return output
}
