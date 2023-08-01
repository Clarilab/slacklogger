package slacklogger

import (
	"fmt"
)

type SlackLogger struct {
	webhookURL  string
	environment string
	isDebug     bool
}

// NewSlackLogger returns a new instance of SlackLogger
func NewSlackLogger(webhookURL, environment string, isDebug bool) *SlackLogger {
	return &SlackLogger{
		webhookURL:  webhookURL,
		environment: environment,
		isDebug:     isDebug,
	}
}

func (logger *SlackLogger) Log(message string) {
	LogWithURL(message, logger.webhookURL, logger.environment, logger.isDebug)
}

func LogWithURL(message, url, env string, isDebug bool) {
	if env != "" {
		message = fmt.Sprintf("env=%s, %s", env, message)
	}

	if isDebug {
		fmt.Println("Debug: Logging to Slack: " + message)

		return
	}

	err := Send(url, Payload{Text: message})
	if err != nil {
		fmt.Printf("Error while logging to Slack: %s\nOriginal message was: %s\n", err, message)
	}
}
