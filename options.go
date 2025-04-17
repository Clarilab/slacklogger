package slacklogger

// Option is an option func for creating a new SlackLogger.
type Option func(*SlackLogger)

// WithAuthorization sets the authorization token.
func WithAuthorization(token string) Option {
	return func(sl *SlackLogger) {
		sl.token = token
	}
}

// UseDebug is an option for creating a new SlackLogger.
// When used the logger is NOT logging to slack, instead its logging to stdout using log/slog.
// Can be useful for tests.
func UseDebug() Option {
	return func(sl *SlackLogger) {
		sl.isDebug = true
	}
}

// WithEnvironment is an option for creating a new SlackLogger.
// When used the given environment is added before the message.
func WithEnvironment(env string) Option {
	return func(sl *SlackLogger) {
		sl.environment = env
	}
}
