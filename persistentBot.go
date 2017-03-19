package main

import (
	"log"

	"github.com/asdine/storm"
)

//PersistentBot extends original Bot struct by adding persistence of requests
//using BoltDB and Storm ORM
type PersistentBot struct {
	*Bot

	db *storm.DB
}

func (bot *PersistentBot) Init() {
	bot.Bot.Init()

	bot.initRequestsFromDb()
	log.Println("Persistent Bot initialization")
}

func (bot *PersistentBot) initRequestsFromDb() {
	var requests []*Request
	err := bot.db.All(&requests)
	if err != nil {
		log.Panic(err)
	}
	bot.Requests = requests
}

func (bot *PersistentBot) AddRequest(requestString string) (string, *Request) {
	result, request := bot.Bot.AddRequest(requestString)

	bot.saveRequestToDb(request)

	return result, request
}

func (bot *PersistentBot) CloseRequest(rawRequestNum string) (string, *Request) {
	result, request := bot.Bot.CloseRequest(rawRequestNum)

	bot.saveRequestToDb(request)

	return result, request
}

func (bot *PersistentBot) saveRequestToDb(request *Request) {
	if bot.db != nil {
		bot.db.Save(request)
	}
}

func (bot *PersistentBot) FinishWork() {
	bot.db.Close()
	log.Println("Persistent Bot finishes it's work")

	bot.Bot.FinishWork()
}

func getPersistentBot() *PersistentBot {
	bot := getBot()

	persistentBot := PersistentBot{Bot: bot, db: initDb(DB_FILENAME)}
	persistentBot.Init()

	return &persistentBot
}

func (bot *PersistentBot) Shutdown() {
	bot.FinishWork()

	bot.Bot.Shutdown()
}