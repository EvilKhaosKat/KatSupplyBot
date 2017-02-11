package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const RequestOne = "request one"
const RequestTwo = "request two"

const RequestOpen = "open request"

const RequestToBeClosed = "request to be closed"

func TestBotRequestsAddOne(t *testing.T) {
	bot := &Bot{}

	bot.AddRequest(RequestOne)

	require.Len(t, bot.Requests, 1)
}

func TestBotRequestsAddMultiple(t *testing.T) {
	bot := &Bot{}

	bot.AddRequest(RequestOne)
	bot.AddRequest(RequestTwo)

	require.Len(t, bot.Requests, 2)
}

func TestBotRequestsToText(t *testing.T) {
	bot := &Bot{}

	bot.AddRequest(RequestOne)
	bot.AddRequest(RequestTwo)

	requestsText := bot.GetRequestsText()
	t.Log("requestsText:", requestsText)

	require.Contains(t, requestsText, RequestOne)
	require.Contains(t, requestsText, RequestTwo)
}

func TestBotCloseRequest(t *testing.T) {
	bot := &Bot{}

	bot.AddRequest(RequestOpen)
	bot.AddRequest(RequestToBeClosed)

	bot.CloseRequest("1") //count from 0

	requestsText := bot.GetRequestsText()
	t.Log("requestsText:", requestsText)

	require.Contains(t, requestsText, RequestOpen)
	require.NotContains(t, requestsText, RequestToBeClosed)
}
