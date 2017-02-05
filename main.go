package main

import (
	"bufio"
	"fmt"
	telegramBotApi "gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
)

const UPDATES_TIMEOUT = 60

const COMMAND_ADD = "add"
const COMMAND_LIST = "list"
const COMMAND_CLOSE = "close"

type Request struct {
	Id     int `storm:"id,increment"`
	Name   string
	Closed bool
}


func main() {
	bot := getBot()

	for update := range bot.getUpdatesChan() {
		if update.Message == nil {
			continue
		}

		handleUpdate(update, bot)
	}

	defer bot.FinishWork()
}

//Logic of messages handling
func handleUpdate(update telegramBotApi.Update, bot *Bot) {
	//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	message := update.Message

	if message.IsCommand() {
		commandArguments := message.CommandArguments()

		switch command := message.Command(); command {
		case COMMAND_ADD:
			result := bot.addRequest(commandArguments)
			bot.sendReply(update, result)
		case COMMAND_LIST:
			bot.sendReply(update, bot.getRequestsText())
		case COMMAND_CLOSE:
			result := bot.closeRequest(commandArguments)
			bot.sendReply(update, result)
		default:
			bot.sendReply(update, fmt.Sprintf("I can't understart command '%s'", command))
		}
	}
}

func (request Request) String() string {
	return request.Name
}

func readTokenFile() string {
	file, err := os.Open("token")
	if err != nil {
		log.Panic("token file reading error:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	return scanner.Text()
}
