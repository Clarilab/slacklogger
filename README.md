# slacklogger

This client is used to send messages using a slack webhook url.

## Installation

```shell
go get github.com/Clarilab/slacklogger/v2
```

## Importing

```go
import "github.com/Clarilab/slacklogger/v2"
```

## Examples

### Logging with an instanced logger

```go
webhookURL := "https://hooks.slack.com/..."
environment := "development"
isDebug := false

slacker := slacklogger.NewSlackLogger(webhookURL, environment, isDebug)

slacker.Log("Something weird")

// this will result in: env=development Something weird
```

### Logging without an instanced logger

```go
webhookURL := "https://hooks.slack.com/..."
environment := ""
isDebug := false
message := "Hello World!"


slacklogger.LogWithURL(message, webhookURL, environment, isDebug)

// this will result in: Hello World!
```

If isDebug is set to true, it will print to stdout instead.
