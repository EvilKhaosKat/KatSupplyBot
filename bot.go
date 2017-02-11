package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	telegramBotApi "gopkg.in/telegram-bot-api.v4"
)

type BotCommunicationInterface interface {
	AddRequest(requestString string) (string, *Request)
	GetRequestsText() string
	CloseRequest(rawRequestNum string) (string, *Request)

	SendReply(update telegramBotApi.Update, text string)
}

//Bot represents Bot entity, associated with telegramBotApi
//After creating an exemplar it's highly recommended to call #Init method
//At the end of work it's highly recommended to call #FinishWork method
type Bot struct {
	Requests []*Request
	botApi   *telegramBotApi.BotAPI
}

func (bot *Bot) getUpdatesChan() <-chan telegramBotApi.Update {
	u := telegramBotApi.NewUpdate(0)

	u.Timeout = UpdatesTimeout

	updates, err := bot.botApi.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	return updates
}

func (bot *Bot) SendReply(update telegramBotApi.Update, text string) {
	replyMessage := telegramBotApi.NewMessage(update.Message.Chat.ID, text)
	replyMessage.ReplyToMessageID = update.Message.MessageID

	bot.botApi.Send(replyMessage)
}

func (bot *Bot) AddRequest(requestString string) (string, *Request) {
	if len(requestString) == 0 {
		return "Empty request won't be added", nil
	}

	request := &Request{Name: requestString}
	bot.Requests = append(bot.Requests, request)

	return fmt.Sprintf("Request '%s' added", request), request
}

func (bot *Bot) GetRequestsText() string {
	if len(bot.Requests) == 0 {
		return "No active requests at the moment"
	}

	var buffer bytes.Buffer

	for number, request := range bot.Requests {
		if !request.Closed {
			buffer.WriteString(fmt.Sprintf("%d: %s\n", number, request))
		}
	}

	return buffer.String()
}

func (bot *Bot) CloseRequest(rawRequestNum string) (string, *Request) {
	if len(rawRequestNum) == 0 {
		return "Request number to close required", nil
	}

	requestNum, err := strconv.Atoi(rawRequestNum)
	if err != nil {
		return fmt.Sprintf("Request number to close required, but got error: %s", err.Error()), nil
	}

	if requestNum < 0 || requestNum > len(bot.Requests) {
		return "Incorrect request number", nil
	}

	request := bot.Requests[requestNum]
	if request.Closed {
		return fmt.Sprintf("Request '%s' is already closed", request), nil
	}

	request.Closed = true

	return fmt.Sprintf("Request '%s' closed", request), request
}

func (bot *Bot) FinishWork() {
	log.Println("Bot finishes it's work")
}

func (bot *Bot) Init() {
	log.Println("Bot initialization")
}

func getBotApi(token string) *telegramBotApi.BotAPI {
	bot, err := telegramBotApi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	return bot
}

func getBot() *Bot {
	log.Println("Trying to read 'token' file")
	token := readTokenFile()
	log.Println("Token acquired")

	botApi := getBotApi(token)
	log.Printf("Authorized on account %s", botApi.Self.UserName)

	bot := &Bot{botApi: botApi}

	return bot
}

func getPersistentBot() *PersistentBot {
	bot := getBot()

	persistentBot := PersistentBot{Bot: bot, db: initDb(DB_FILENAME)}
	persistentBot.Init()

	return &persistentBot
}
