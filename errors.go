package slacklogger

func newSlackError(status string) error {
	return &SlackError{status: status}
}

// SlackError occurs when slack responded with an error status code.
type SlackError struct {
	status string
}

// Error implements the error interface.
func (e *SlackError) Error() string {
	return "error sending to slack. status: " + e.status
}
