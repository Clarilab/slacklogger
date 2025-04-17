package slacklogger

import (
	"context"
	"fmt"
	"net/http"

	"log/slog"
)

// SlackLogger provides functionality to send messages to slack.
type SlackLogger struct {
	client      *http.Client
	url         string
	token       string
	channel     string
	environment string
	isDebug     bool
}

// NewSlackLogger creates a new instance of SlackLogger.
func NewSlackLogger(url, channel string, options ...Option) *SlackLogger {
	sl := SlackLogger{
		client:  new(http.Client),
		url:     url,
		channel: channel,
	}

	for i := range options {
		options[i](&sl)
	}

	return &sl
}

// Log sends a simple message to slack.
func (l *SlackLogger) Log(message string) {
	l.log(context.Background(), message)
}

// LogContext sends a simple message to slack with the given context.Context.
func (l *SlackLogger) LogContext(ctx context.Context, message string) {
	l.log(ctx, message)
}

// Write implements the io.Writer interface.
func (l *SlackLogger) Write(message []byte) (int, error) {
	l.Log(string(message))

	return len(message), nil
}

// Send sends the given payload to slack.
func (l *SlackLogger) Send(ctx context.Context, payload *Payload) {
	const errMessage = "error while logging to slack"

	if err := l.send(ctx, payload); err != nil {
		slog.Error(errMessage, ErrorAttr(err))
	}
}

func (l *SlackLogger) log(ctx context.Context, message string) {
	const (
		errMessage      = "error while logging to slack"
		envPrefixFormat = "environment: %s\n\n, %s"
		debugMessage    = "pretending to log to slack"
	)

	if l.environment != "" {
		message = fmt.Sprintf(envPrefixFormat, l.environment, message)
	}

	if l.isDebug {
		slog.Debug(debugMessage, MessageAttr(message))

		return
	}

	payload := Payload{
		Channel: l.channel,
		Text:    message,
	}

	if err := l.send(ctx, &payload); err != nil {
		slog.Error(errMessage, ErrorAttr(err))
	}
}
