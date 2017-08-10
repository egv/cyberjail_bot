package main

import (
	"log"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Any is predicate method for strings
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func main() {
	var blackList = []string{"darknet", "даркнет"}
	token := os.Getenv("TELEGRAMBOTTOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		var message = update.Message
		var messageText = strings.ToLower(message.Text)

		if message.ForwardFromChat != nil && message.ForwardFromChat.ID != message.Chat.ID && Any(blackList, func(word string) bool {
			return strings.Contains(messageText, word)
		}) {
			config := tgbotapi.ChatMemberConfig{
				ChatID: message.Chat.ID,
				UserID: message.From.ID,
			}
			kickConfig := tgbotapi.KickChatMemberConfig{
				ChatMemberConfig: config,
			}
			bot.KickChatMember(kickConfig)
		}

		replies := map[string]string{
			"((":     "Чо ты такой грустный?",
			"гук":    "В академии наук заседает чёртов гук. Джонни, они научились прятаться даже там!",
			"shrug":  "¯\\_(ツ)_/¯",
			"дорого": "Твоя нищета омерзительна",
			"фашист": "По поводу взаимоотношения фашизма и расизма в науке существуют разные мнения. Сторонники одной точки зрения полагают, что идея биологического расизма была прерогативой нацистского режима, тогда как в фашизме упор делается на нацию, а не расу. Последователи этой теории в целом склонны выделять нацизм как особый исторический феномен, не считая его одной из разновидностей фашизма",
		}

		for trigger, reply := range replies {
			if strings.Contains(messageText, trigger) {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}

		stickerReplies := map[string]string{
			"гитлер": "CAADAgADuQMAApj8Lgej7SMzfLsfmwI",
			"пиздец": "CAADAQADpQMAAjbzwAuEVbJbdhMgRAI",
			"хахаха": "CAADAgADCwEAAvR7GQABuArOzKHFjusC",
			"хммм":   "CAADAgADiQMAAjbsGwVkQ0mbOVmPTQI",
		}

		for trigger, reply := range stickerReplies {
			if strings.Contains(messageText, trigger) {
				msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, reply)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}
}
