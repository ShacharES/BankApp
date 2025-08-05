package token

import "time"

// Maker is an interface that defines methods for creating and verifying tokens.
type Maker interface {
	// CreateToken creates a new token for a specific user ID with a given duration.
	CreateToken(userID string, duration time.Duration) (string, error)

	// VerifyToken checks the validity of a token and returns its payload if valid.
	VerifyToken(token string) (*Payload, error)
}
