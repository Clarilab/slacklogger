# slacklogger

This lib can be used to easily send messages to slack channels via the Slack Web API.

## Installation

```shell
go get github.com/Clarilab/slacklogger/v3
```

## Importing

```go
import "github.com/Clarilab/slacklogger/v3"
```

## Requirements
A [**Slack-App**](https://api.slack.com/docs/apps) installed to the desired **Slack-Workspace** with correct permissions/scopes set.

Minimum required permission/scope is: **chat:write**

For more information checkout the [Slack Apps Quickstart Guide](https://api.slack.com/quickstart).

## Authorization
When using a proxy like [**Slack-Proxy**](https://github.com/fortio/slack-proxy), the authorization needs to be setup in the proxy and is not needed here.

Otherwise the [**Slack-Apps**](https://api.slack.com/docs/apps) **OAuth-Token** needs to be provided via the WithAuthorization() option.

## Examples

```go
url := "https://<workspace-name>.slack.com/api/chat.postMessage"
token := "slack-token"
channel := "log-channel"
environment := "development"



slacker := slacklogger.NewSlackLogger(url, channel, slacklogger.WithEnvironment(environment), slacklogger.WithAuthorization(token))

slacker.Log("Something weird")

// this will result in: 
// environment: development 
// 
// Something weird
```

If the UseDebug option is used, it will log to stdout instead using the log/slog logger.
