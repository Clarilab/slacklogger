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

	// LogWithName prints a simple text to slack, while using the given url as webhookURL
	LogWithURL(url, text string)

	// LogWithURLAndName prints a formated text, containing the name, to slack, while using the given url as webhookURL
	LogWithURLAndName(url, name, text string)
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
	message := fmt.Sprintf("env=%s: %s", logger.environment, text)

	if !logger.onlyProd || !logger.isProd {
		fmt.Println("Logging to Slack: " + message)
	}

	payload := slack.Payload{
		Text:     message,
		Username: logger.username,
		Channel:  logger.channel,
	}

	err := slack.Send(logger.webhookURL, "", payload)
	if len(err) > 0 {
		fmt.Printf("Error while logging to Slack: %s\nOriginal message was: %s\n", err, message)
	}
}

func (logger *slackLogger) LogWithName(name, text string) {
	message := fmt.Sprintf("%s: %s", name, text)
	logger.Log(message)
}

func (logger *slackLogger) LogWithURL(url, text string) {
	message := fmt.Sprintf("env=%s: %s", logger.environment, text)

	if logger.onlyProd && !logger.isProd {
		fmt.Println("Logging to Slack: " + message)

		return
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

func (logger *slackLogger) LogWithURLAndName(url, name, text string) {
	message := fmt.Sprintf("%s: %s", name, text)
	logger.LogWithURL(url, message)
}
