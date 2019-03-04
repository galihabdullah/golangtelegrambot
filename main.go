package main


import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("653629960:AAFizrqn043ird_bXJpMZmkDCtpSjyWfmDA")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
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