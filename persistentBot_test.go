package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const TestDbAddOneFilename = "TestPersistentBotRequestsAddOne_test.db"
const TestDbAddMultipleFilename = "TestPersistentBotRequestsAddMultiple_test.db"
const TestDbRequestsToTextFilename = "TestPersistentBotRequestsToText_test.db"
const TestDbCloseRequestFilename = "TestPersistentBotCloseRequest_test.db"

func getTestPersistentBot(testDbFilename string) *PersistentBot {
	bot := PersistentBot{Bot: &Bot{}, db: initDb(testDbFilename)}
	bot.Init()

	return &bot
}

//runs tests, and removes test db files
func TestMain(m *testing.M) {
	code := m.Run()
	removeTestDbFiles()
	os.Exit(code)
}

func removeTestDbFiles() {
	os.Remove(TestDbAddOneFilename)
	os.Remove(TestDbAddMultipleFilename)
	os.Remove(TestDbRequestsToTextFilename)
	os.Remove(TestDbCloseRequestFilename)
}

func TestPersistentBotRequestsAddOne(t *testing.T) {
	bot := getTestPersistentBot(TestDbAddOneFilename)
	bot.AddRequest(RequestOne)
	bot.FinishWork()

	bot = getTestPersistentBot(TestDbAddOneFilename)

	require.Len(t, bot.Requests, 1)
}

func TestPersistentBotRequestsAddMultiple(t *testing.T) {
	bot := getTestPersistentBot(TestDbAddMultipleFilename)
	bot.AddRequest(RequestOne)
	bot.AddRequest(RequestTwo)
	bot.FinishWork()

	bot = getTestPersistentBot(TestDbAddMultipleFilename)

	require.Len(t, bot.Requests, 2)
}

func TestPersistentBotRequestsToText(t *testing.T) {
	bot := getTestPersistentBot(TestDbRequestsToTextFilename)
	bot.AddRequest(RequestOne)
	bot.AddRequest(RequestTwo)
	bot.FinishWork()

	bot = getTestPersistentBot(TestDbRequestsToTextFilename)

	requestsText := bot.GetRequestsText()
	t.Log("requestsText:", requestsText)

	require.Contains(t, requestsText, RequestOne)
	require.Contains(t, requestsText, RequestTwo)
}

func TestPersistentBotCloseRequest(t *testing.T) {
	bot := getTestPersistentBot(TestDbCloseRequestFilename)
	bot.AddRequest(RequestOpen)
	bot.AddRequest(RequestToBeClosed)
	bot.FinishWork()

	bot = getTestPersistentBot(TestDbCloseRequestFilename)

	bot.CloseRequest("1") //count from 0

	requestsText := bot.GetRequestsText()
	t.Log("requestsText:", requestsText)

	require.Contains(t, requestsText, RequestOpen)
	require.NotContains(t, requestsText, RequestToBeClosed)
}
