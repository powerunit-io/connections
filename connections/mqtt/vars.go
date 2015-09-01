package mqtt

var (
	// AvailableConnectionTypes -
	AvailableConnectionTypes = []string{"tcp", "tls", "ws"}

	// InitialConnectionTimeout -
	InitialConnectionTimeout = 10

	// MaxTopicSubscribeAttempts -
	MaxTopicSubscribeAttempts = 5
)
