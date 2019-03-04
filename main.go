package main

import (
	"io/ioutil"
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"os/exec"
	"strings"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("653629960:AAFizrqn043ird_bXJpMZmkDCtpSjyWfmDA")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if strings.ContainsAny(update.Message.Text, "func"){
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			message := []byte(update.Message.Text)
			ioutil.WriteFile("livetest.go", message, os.ModePerm)
			out, err := exec.Command("go", "run", "livetest.go").Output()
			if err != nil{
				msg.Text = err.Error()
			}else{
				msg.Text = "outputnya adalah " + string(out)
			}
			bot.Send(msg)
			log.Printf("To: %+v Text: %+v\n",msg.ReplyToMessageID, msg.Text)
		}
	}
}