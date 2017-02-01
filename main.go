package main

import (
	"bufio"
	"bytes"
	"fmt"
	telegramBotApi "gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
	"strconv"
)

const UPDATES_TIMEOUT = 60

const COMMAND_ADD = "add"
const COMMAND_LIST = "list"
const COMMAND_CLOSE = "close"

var botApi *telegramBotApi.BotAPI

type Request struct {
	name   string
	closed bool
}

type Bot struct {
	requests []*Request
	botApi   *telegramBotApi.BotAPI
}

func main() {
	bot := getBot()

	for update := range bot.getUpdatesChan() {
		if update.Message == nil {
			continue
		}

		handleUpdate(update, bot)
	}
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
	return request.name
}

func (bot *Bot) closeRequest(rawRequestNum string) string {
	requestNum, err := strconv.Atoi(rawRequestNum)
	if err != nil {
		return err.Error()
	}

	if requestNum < 0 || requestNum > len(bot.requests) {
		return "Incorrect request number"
	}

	request := bot.requests[requestNum]
	if request.closed {
		return fmt.Sprintf("Request '%s' is already closed", request)
	}

	request.closed = true
	//bot.requests = append(bot.requests[:requestNum], bot.requests[requestNum+1])
	return fmt.Sprintf("Request '%s' closed", request)
}

func (bot *Bot) getRequestsText() string {
	if len(bot.requests) == 0 {
		return "No active requests at the moment"
	}

	var buffer bytes.Buffer

	for number, request := range bot.requests {
		if !request.closed {
			buffer.WriteString(fmt.Sprintf("%d: %s\n", number, request))
		}
	}

	return buffer.String()
}

func getBot() *Bot {
	log.Println("Trying to read 'token' file")
	token := readTokenFile()
	log.Println("Token acquired")

	botApi = getBotApi(token)
	log.Printf("Authorized on account %s", botApi.Self.UserName)

	bot := &Bot{botApi: botApi}
	return bot
}

func (bot *Bot) sendReply(update telegramBotApi.Update, text string) {
	replyMessage := telegramBotApi.NewMessage(update.Message.Chat.ID, text)
	replyMessage.ReplyToMessageID = update.Message.MessageID

	bot.botApi.Send(replyMessage)
}

func (bot *Bot) getUpdatesChan() <-chan telegramBotApi.Update {
	u := telegramBotApi.NewUpdate(0)

	u.Timeout = UPDATES_TIMEOUT

	updates, err := bot.botApi.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	return updates
}

func (bot *Bot) addRequest(name string) string {
	request := &Request{name: name}
	bot.requests = append(bot.requests, request)

	return fmt.Sprintf("Request '%s' added", name)
}

func getBotApi(token string) *telegramBotApi.BotAPI {
	bot, err := telegramBotApi.NewBotAPI(token)
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
