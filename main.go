package main

import (
	"log"
	"os"
	"bufio"
	"fmt"
	telegramBot "gopkg.in/telegram-bot-api.v4"
)

const UPDATES_TIMEOUT = 60

const COMMAND_ADD = "add"
const COMMAND_LIST = "list"
const COMMAND_CLOSE = "close"


func main() {
	log.Println("Trying to read 'token' file")
	token := readTokenFile()
	log.Println("Token acquired")

	bot := getBot(token)
	log.Printf("Authorized on account %s", bot.Self.UserName)

	for update := range getUpdatesChan(bot) {
		if update.Message == nil {
			continue
		}

		handleUpdate(update, bot)
	}
}

//Logic of messages handling
func handleUpdate(update telegramBot.Update, bot *telegramBot.BotAPI) {
	//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	message := update.Message

	if message.IsCommand() {
		//commandArguments := message.CommandArguments()

		switch command := message.Command(); command {
		case COMMAND_ADD:

		case COMMAND_LIST:

		case COMMAND_CLOSE:

		default:
			sendReply(update, bot,
				fmt.Sprintf("I can't understart command '%s'", command))
		}
	}
}

func sendReply(update telegramBot.Update, bot *telegramBot.BotAPI, text string) {
	replyMessage := telegramBot.NewMessage(update.Message.Chat.ID, text)
	replyMessage.ReplyToMessageID = update.Message.MessageID

	bot.Send(replyMessage)
}

func getUpdatesChan(bot *telegramBot.BotAPI) <- chan telegramBot.Update {
	u := telegramBot.NewUpdate(0)

	u.Timeout = UPDATES_TIMEOUT

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	return updates
}

func getBot(token string) *telegramBot.BotAPI {
	bot, err := telegramBot.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	return bot
}

func readTokenFile() string {
	file, err := os.Open("token")
	if err != nil {
		fmt.Println("token file reading error:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	return scanner.Text()
}
