package protocol

import "strings"

// ProcessMessage processes a message according to the custom protocol.
func ProcessMessage(message string) string {
	message = strings.TrimSpace(message)
	if message == "HELLO" {
		return "WORLD"
	}
	return "UNKNOWN COMMAND"
}
