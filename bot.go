package main

import (
	telegramBotApi "gopkg.in/telegram-bot-api.v4"
	"log"
	"fmt"
	"bytes"
	"strconv"
)

type Bot struct {
	requests []*Request
	botApi   *telegramBotApi.BotAPI
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

func (bot *Bot) sendReply(update telegramBotApi.Update, text string) {
	replyMessage := telegramBotApi.NewMessage(update.Message.Chat.ID, text)
	replyMessage.ReplyToMessageID = update.Message.MessageID

	bot.botApi.Send(replyMessage)
}

func (bot *Bot) addRequest(requestString string) string {
	if len(requestString) == 0 {
		return "Empty request won't be added"
	}

	request := &Request{name: requestString}
	bot.requests = append(bot.requests, request)

	return fmt.Sprintf("Request '%s' added", request)
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

func (bot *Bot) closeRequest(rawRequestNum string) string {
	if len(rawRequestNum) == 0 {
		return "Request number to close required"
	}

	requestNum, err := strconv.Atoi(rawRequestNum)
	if err != nil {
		return fmt.Sprintf("Request number to close required, but got error: %s", err.Error())
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

func getBot() *Bot {
	log.Println("Trying to read 'token' file")
	token := readTokenFile()
	log.Println("Token acquired")

	botApi := getBotApi(token)
	log.Printf("Authorized on account %s", botApi.Self.UserName)

	bot := &Bot{botApi: botApi}
	return bot
}