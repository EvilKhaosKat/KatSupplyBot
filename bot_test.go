package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

const REQUEST_ONE = "request one"
const REQUEST_TWO = "request two"

const REQUEST_OPEN = "open request"

const REQUEST_TO_BE_CLOSED = "request to be closed"

func TestBotRequestsAddOne(t *testing.T) {
	bot := &Bot{}

	bot.addRequest(REQUEST_ONE)

	require.Len(t, bot.requests, 1)
}

func TestBotRequestsAddMultiple(t *testing.T) {
	bot := &Bot{}

	bot.addRequest(REQUEST_ONE)
	bot.addRequest(REQUEST_TWO)

	require.Len(t, bot.requests, 2)
}

func TestBotRequestsToText(t *testing.T) {
	bot := &Bot{}

	bot.addRequest(REQUEST_ONE)
	bot.addRequest(REQUEST_TWO)

	requestsText := bot.getRequestsText()
	t.Log("requestsText:", requestsText)

	require.Contains(t, requestsText, REQUEST_ONE)
	require.Contains(t, requestsText, REQUEST_TWO)
}

func TestBotCloseRequest(t *testing.T) {
	bot := &Bot{}

	bot.addRequest(REQUEST_OPEN)
	bot.addRequest(REQUEST_TO_BE_CLOSED)

	bot.closeRequest("1") //count from 0

	requestsText := bot.getRequestsText()
	t.Log("requestsText:", requestsText)

	require.Contains(t, requestsText, REQUEST_OPEN)
	require.NotContains(t, requestsText, REQUEST_TO_BE_CLOSED)
}
