package slacklogger_test

import (
	"os"
	"testing"

	"github.com/Clarilab/slacklogger/v2"
)

func Test_Log(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		t.Fatal("webhook url is not set")
	}

	logger := slacklogger.NewSlackLogger(webhookURL, "dev", false)
	logger.Log("test message")
}
