package main

import (
	"github.com/ashwanthkumar/slack-go-webhook"
)

const url = "https://hooks.slack.com/services/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"

func SlackMsg(msg string) {
	p := slack.Payload{
		Text: msg,
	}

	slack.Send(url, "", p)
}
