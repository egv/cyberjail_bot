package main

import (
	"fmt"
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

			kickedUser := message.From
			kickMessage := fmt.Sprintf("Баним %s %s (@%s)", kickedUser.FirstName, kickedUser.LastName, kickedUser.UserName)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, kickMessage)

			messageToDelete := tgbotapi.DeleteMessageConfig{
				ChatID:    message.Chat.ID,
				MessageID: message.MessageID,
			}

			bot.DeleteMessage(messageToDelete)
			bot.Send(msg)
			bot.KickChatMember(kickConfig)
		}
	}
}
