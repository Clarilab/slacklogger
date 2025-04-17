package slacklogger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Payload is the payload send to slack.
type Payload struct {
	Channel   string          `json:"channel"` // required
	Text      string          `json:"text"`
	AsUser    bool            `json:"as_user"`
	Username  string          `json:"username,omitempty"`
	IconURL   string          `json:"icon_url,omitempty"`
	IconEmoji string          `json:"icon_emoji,omitempty"`
	ThreadTS  string          `json:"thread_ts,omitempty"`
	Parse     string          `json:"parse,omitempty"`
	LinkNames bool            `json:"link_names,omitempty"`
	Blocks    json.RawMessage `json:"blocks,omitempty"` // JSON serialized array of blocks
}

func (l *SlackLogger) send(ctx context.Context, payload *Payload) error {
	const (
		errMessage          = "failed to send to slack: %w"
		headerContentType   = "Content-Type"
		headerAuthorization = "Authorization"
		mimeJSON            = "application/json; charset=utf-8"
		tokenPrefix         = "Bearer "
	)

	payloadBytes, err := json.Marshal(&payload)
	if err != nil {
		return fmt.Errorf(errMessage, err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, l.url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf(errMessage, err)
	}

	req.Header.Set(headerContentType, mimeJSON)

	if l.token != "" {
		req.Header.Set(headerAuthorization, tokenPrefix+l.token)
	}

	resp, err := l.client.Do(req)
	if err != nil {
		return fmt.Errorf(errMessage, err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf(errMessage, newSlackError(resp.Status))
	}

	return nil
}
