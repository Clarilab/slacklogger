package slacklogger_test

import (
	"os"
	"testing"

	"github.com/Clarilab/slacklogger/v3"
)

func Test_Log(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	url := os.Getenv("SLACK_URL")
	if url == "" {
		t.Fatal("slack url is not set")
	}

	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		t.Fatal("slack token is not set")
	}

	logger := slacklogger.NewSlackLogger(url, "logs-test", slacklogger.WithAuthorization(token))
	logger.Log("test message")
}
