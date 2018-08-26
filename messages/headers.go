package messages

import (
	"time"

	"github.com/resc/gobit/claims"
)

type (
	// Headers contains the message metadata
	Headers interface {
		// App is the principal representing the application that created the message
		App() claims.Principal
		CorrelationId() string
		MessageId() string
		ReplyTo() string
		Timestamp() time.Time
		Type() string
		// User is the principal representing the user authenticated by App() that created the message.
		// It's the same as App() if there was no humanin the loop
		User() claims.Principal
	}
)
