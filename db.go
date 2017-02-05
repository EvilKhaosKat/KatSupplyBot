package main

import (
	"github.com/asdine/storm"
	"log"
)

const DB_FILENAME = "KatSupplyBot.db"

func initDb() *storm.DB {
	db, err := storm.Open(DB_FILENAME)

	if err != nil {
		log.Panic(err)
	}

	return db
}
