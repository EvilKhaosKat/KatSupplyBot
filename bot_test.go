package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const REQUEST_ONE = "request one"
const REQUEST_TWO = "request two"

const REQUEST_OPEN = "open request"

const REQUEST_TO_BE_CLOSED = "request to be closed"

func TestBotRequestsAddOne(t *testing.T) {
	bot := &Bot{}

	bot.AddRequest(REQUEST_ONE)

	require.Len(t, bot.Requests, 1)
}

func TestBotRequestsAddMultiple(t *testing.T) {
	bot := &Bot{}

	bot.AddRequest(REQUEST_ONE)
	bot.AddRequest(REQUEST_TWO)

	require.Len(t, bot.Requests, 2)
}

func TestBotRequestsToText(t *testing.T) {
	bot := &Bot{}

	bot.AddRequest(REQUEST_ONE)
	bot.AddRequest(REQUEST_TWO)

	requestsText := bot.GetRequestsText()
	t.Log("requestsText:", requestsText)

	require.Contains(t, requestsText, REQUEST_ONE)
	require.Contains(t, requestsText, REQUEST_TWO)
}

func TestBotCloseRequest(t *testing.T) {
	bot := &Bot{}

	bot.AddRequest(REQUEST_OPEN)
	bot.AddRequest(REQUEST_TO_BE_CLOSED)

	bot.CloseRequest("1") //count from 0

	requestsText := bot.GetRequestsText()
	t.Log("requestsText:", requestsText)

	require.Contains(t, requestsText, REQUEST_OPEN)
	require.NotContains(t, requestsText, REQUEST_TO_BE_CLOSED)
}
