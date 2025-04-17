package slacklogger

import "log/slog"

const (
	keyError   = "error"
	keyMessage = "message"
)

// ErrorAttr creates an slog.Attr for error.
func ErrorAttr(err error) slog.Attr {
	return slog.Any(keyError, err)
}

// MessageAttr creates an slog.Attr for message.
func MessageAttr(msg string) slog.Attr {
	return slog.String(keyMessage, msg)
}
