package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	telegramBotApi "gopkg.in/telegram-bot-api.v4"
	"strings"
)

//UpdatesTimeout Telegram API poll timeout
const UpdatesTimeout = 60

const CommandAdd = "add"
const CommandList = "list"
const CommandClose = "close"
const CommandShutdown = "shutdown"

type Request struct {
	ID     int `storm:"id,increment"`
	Name   string
	Closed bool
}

func main() {
	bot := getPersistentBot()

	for update := range bot.getUpdatesChan() {
		if update.Message == nil {
			continue
		}

		handleUpdate(update, bot)
	}

	defer bot.FinishWork()
}

//Logic of messages handling
func handleUpdate(update telegramBotApi.Update, bot BotCommunicationInterface) {
	//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	message := update.Message

	if message.IsCommand() {
		commandArguments := message.CommandArguments()

		switch command := strings.ToLower(message.Command()); command {
		case CommandAdd:
			result, _ := bot.AddRequest(commandArguments)
			bot.SendReply(update, result)
		case CommandList:
			bot.SendReply(update, bot.GetRequestsText())
		case CommandClose:
			result, _ := bot.CloseRequest(commandArguments)
			bot.SendReply(update, result)
		case CommandShutdown:
			username := message.From.UserName
			if bot.IsAdmin(username) {
				log.Printf("Shutdown initiated by '%s' ", username)

				bot.SendReply(update, "Shutdown initiated")
				bot.Shutdown()
			} else {
				bot.SendReply(update, "You are not authorized to perform shutdown command")
			}
		default:
			bot.SendReply(update, fmt.Sprintf("I can't understand command '%s'", command))
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
