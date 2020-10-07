package slacklogger

import (
	"fmt"

	"github.com/ashwanthkumar/slack-go-webhook"
)

// SlackLogger holds all callable methods
type SlackLogger interface {
	// Log prints a simple text to slack
	Log(text string)

	// LogWithName prints a formated text, containing the name, to slack
	LogWithName(name, text string)

	// LogWithChannel prints a text to a given slack channel
	LogWithChannel(channel, text string)

	// LogWithName prints a simple text to slack, while using the given url as webhookURL
	LogWithURL(url, text string)

	// LogWithURLAndName prints a formated text, containing the name, to slack, while using the given url as webhookURL
	LogWithURLAndName(url, name, text string)

	// LogWithName prints a simple text to a given slack channel, while using the given url as webhookURL
	LogWithChannelAndURL(channel, url, text string)

	// LogWithURLAndName prints a formated text, containing the name, to a given slack channel
	LogWithChannelAndName(channel, name, text string)

	// LogWithURLAndName prints a formated text, containing the name, to a given slack channel, while using the given url as webhookURL
	LogWithChannelAndURLAndName(channel, url, name, text string)
}

type slackLogger struct {
	webhookURL  string
	channel     string
	username    string
	environment string
	isProd      bool
	onlyProd    bool
}

// NewSlackLogger returns a new instance of SlackLogger
func NewSlackLogger(webhookURL, channel, username, environment string, isProd, onlyProd bool) SlackLogger {
	return &slackLogger{
		webhookURL:  webhookURL,
		channel:     channel,
		username:    username,
		environment: environment,
		isProd:      isProd,
		onlyProd:    onlyProd,
	}
}

func (logger *slackLogger) Log(text string) {
	logger.LogWithURL(logger.webhookURL, text)
}

func (logger *slackLogger) LogWithChannel(channel, text string) {
	logger.LogWithChannelAndURL(channel, logger.webhookURL, text)
}

func (logger *slackLogger) LogWithURL(url, text string) {
	logger.LogWithChannelAndURL(logger.channel, url, text)
}

func (logger *slackLogger) LogWithName(name, text string) {
	message := fmt.Sprintf("%s: %s", name, text)
	logger.Log(message)
}

func (logger *slackLogger) LogWithURLAndName(url, name, text string) {
	message := fmt.Sprintf("%s: %s", name, text)
	logger.LogWithChannelAndURL(logger.channel, url, message)
}

func (logger *slackLogger) LogWithChannelAndName(channel, name, text string) {
	message := fmt.Sprintf("%s: %s", name, text)
	logger.LogWithChannelAndURL(channel, logger.webhookURL, message)
}

func (logger *slackLogger) LogWithChannelAndURLAndName(channel, url, name, text string) {
	message := fmt.Sprintf("%s: %s", name, text)
	logger.LogWithChannelAndURL(channel, url, message)
}

func (logger *slackLogger) LogWithChannelAndURL(channel, url, text string) {
	message := fmt.Sprintf("env=%s: %s", logger.environment, text)

	if !logger.onlyProd || !logger.isProd {
		fmt.Println("Logging to Slack: " + message)
	}

	payload := slack.Payload{
		Text:     message,
		Username: logger.username,
		Channel:  logger.channel,
	}

	err := slack.Send(url, "", payload)
	if len(err) > 0 {
		fmt.Printf("Error while logging to Slack: %s\nOriginal message was: %s\n", err, message)
	}
}
