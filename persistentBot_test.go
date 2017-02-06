package main

import (
	"testing"
	"github.com/stretchr/testify/require"
	"os"
)

const TEST_DB_ADD_ONE_FILENAME = "TestPersistentBotRequestsAddOne_test.db"
const TEST_DB_ADD_MULTIPLE_FILENAME = "TestPersistentBotRequestsAddMultiple_test.db"
const TEST_DB_REQUESTS_TO_TEXT_FILENAME = "TestPersistentBotRequestsToText_test.db"
const TEST_DB_CLOSE_REQUEST_FILENAME = "TestPersistentBotCloseRequest_test.db"


func getTestPersistentBot(testDbFilename string) *PersistentBot {
	bot := PersistentBot{Bot:&Bot{}, db: initDb(testDbFilename)}
	bot.init()

	return &bot
}

//runs tests, and removes test db files
func TestMain(m *testing.M) {
	code := m.Run()
	removeTestDbFiles()
	os.Exit(code)
}

func removeTestDbFiles() {
	os.Remove(TEST_DB_ADD_ONE_FILENAME)
	os.Remove(TEST_DB_ADD_MULTIPLE_FILENAME)
	os.Remove(TEST_DB_REQUESTS_TO_TEXT_FILENAME)
	os.Remove(TEST_DB_CLOSE_REQUEST_FILENAME)
}

func TestPersistentBotRequestsAddOne(t *testing.T) {
	bot := getTestPersistentBot(TEST_DB_ADD_ONE_FILENAME)
	bot.AddRequest(REQUEST_ONE)
	bot.FinishWork()

	bot = getTestPersistentBot(TEST_DB_ADD_ONE_FILENAME)

	require.Len(t, bot.requests, 1)
}

func TestPersistentBotRequestsAddMultiple(t *testing.T) {
	bot := getTestPersistentBot(TEST_DB_ADD_MULTIPLE_FILENAME)
	bot.AddRequest(REQUEST_ONE)
	bot.AddRequest(REQUEST_TWO)
	bot.FinishWork()

	bot = getTestPersistentBot(TEST_DB_ADD_MULTIPLE_FILENAME)

	require.Len(t, bot.requests, 2)
}

func TestPersistentBotRequestsToText(t *testing.T) {
	bot := getTestPersistentBot(TEST_DB_REQUESTS_TO_TEXT_FILENAME)
	bot.AddRequest(REQUEST_ONE)
	bot.AddRequest(REQUEST_TWO)
	bot.FinishWork()

	bot = getTestPersistentBot(TEST_DB_REQUESTS_TO_TEXT_FILENAME)

	requestsText := bot.GetRequestsText()
	t.Log("requestsText:", requestsText)

	require.Contains(t, requestsText, REQUEST_ONE)
	require.Contains(t, requestsText, REQUEST_TWO)
}

func TestPersistentBotCloseRequest(t *testing.T) {
	bot := getTestPersistentBot(TEST_DB_CLOSE_REQUEST_FILENAME)
	bot.AddRequest(REQUEST_OPEN)
	bot.AddRequest(REQUEST_TO_BE_CLOSED)
	bot.FinishWork()

	bot = getTestPersistentBot(TEST_DB_CLOSE_REQUEST_FILENAME)

	bot.CloseRequest("1") //count from 0

	requestsText := bot.GetRequestsText()
	t.Log("requestsText:", requestsText)

	require.Contains(t, requestsText, REQUEST_OPEN)
	require.NotContains(t, requestsText, REQUEST_TO_BE_CLOSED)
}
