package main

import (
	"github.com/asdine/storm"
	"log"
)

type PersistentBot struct {
	*Bot

	db *storm.DB
}

func (bot *PersistentBot) init() {
	bot.Bot.init()

	bot.initRequestsFromDb()
	log.Println("Persistent Bot initialization")
}

func (bot *PersistentBot) initRequestsFromDb() {
	var requests []*Request
	err := bot.db.All(&requests)
	if err != nil {
		log.Panic(err)
	}
	bot.requests = requests
}

func (bot *PersistentBot) addRequest(requestString string) (string, *Request) {
	result, request := bot.Bot.addRequest(requestString)

	bot.saveRequestToDb(request)

	return result, request
}

func (bot *PersistentBot) closeRequest(rawRequestNum string) (string, *Request) {
	result, request := bot.Bot.closeRequest(rawRequestNum)

	bot.saveRequestToDb(request)

	return result, request
}

func (bot *PersistentBot) saveRequestToDb(request *Request) {
	if bot.db != nil {
		bot.db.Save(request)
	}
}

func (bot *PersistentBot) FinishWork() {
	bot.Bot.FinishWork()

	bot.db.Close()
	log.Println("Persistent Bot finishes it's work")
}
