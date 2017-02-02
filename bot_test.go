package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestBotRequestAdd(t *testing.T) {
	bot := &Bot{}

	//test one
	requestOne := "request one"
	bot.addRequest(requestOne)

	require.Len(t, bot.requests, 1)

	//test two
	requestTwo := "request two"
	bot.addRequest(requestTwo)

	require.Len(t, bot.requests, 2)

	//test three
	requestsText := bot.getRequestsText()
	t.Log("requestsText:", requestsText)
	require.Contains(t, requestsText, requestOne)
	require.Contains(t, requestsText, requestTwo)
}

func TestBotCloseRequest(t *testing.T) {
	bot := &Bot{}

	requestOpen := "open request"
	bot.addRequest(requestOpen)

	requestToBeClosed := "request to be closed"
	bot.addRequest(requestToBeClosed)

	requestNumberToBeClosed := "1" //count from 0
	bot.closeRequest(requestNumberToBeClosed)

	requestsText := bot.getRequestsText()
	t.Log("requestsText:", requestsText)

	//check
	require.Contains(t, requestsText, requestOpen)
	require.NotContains(t, requestsText, requestToBeClosed)
}
