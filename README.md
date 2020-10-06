# slacklogger

This client uses [slack-to-webhook](github.com/ashwanthkumar/slack-go-webhook).

## Installation
```shell
go get github.com/Clarilab/slacklogger
```

## Importing
```go
import "github.com/Clarilab/gocloaksession"
```

## Features
```go
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
```

## Example
```go
webhookURL = "https://hooks.slack.com/..."
channel = "Nicenstein"
username = "Botty McBootface"
environment = "prod"
isProd = true
onlyProd = true

slackLogger := NewSlackLogger(webhookURL, channel, username, environment, isProd, onlyProd)

slackLogger.Log("Something weird")
```