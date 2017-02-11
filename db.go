package main

import (
	"log"

	"github.com/asdine/storm"
)

const DB_FILENAME = "KatSupplyBot.db"

func initDb(dbFilename string) *storm.DB {
	db, err := storm.Open(dbFilename)

	if err != nil {
		log.Panic(err)
	}

	return db
}
