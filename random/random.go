package random

import "github.com/satori/go.uuid"

// RandQueue generates a random queue UUID prepended
func RandQueue() string {
	randUUID := uuid.NewV4().String()
	randUUID = "queue-" + randUUID
	return randUUID
}
