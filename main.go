package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)


func main() {
	bot, err := tgbotapi.NewBotAPI("653629960:AAGVv9s_bq53-qfCcDQ7v_btxbcRwJ1LZD8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = "type /sayhi or /status."
			case "sayhi":
				msg.Text = "Hi :)"
			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}else if strings.ContainsAny(update.Message.Text, "func"){
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			message := []byte(update.Message.Text)
			ioutil.WriteFile("livetest.go", message, os.ModePerm)
			out, err := exec.Command("go", "run", "livetest.go").Output()
			if err != nil{
				msg.Text = "sintax error"
			}else{
				msg.Text = "outputnya adalah " + string(out)
			}
			bot.Send(msg)
		}
		}
	}