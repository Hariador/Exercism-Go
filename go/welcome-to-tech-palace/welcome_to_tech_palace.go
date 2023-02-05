package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	message := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	outPut := strings.Repeat("*", numStarsPerLine) + "\n"
	outPut += welcomeMsg + "\n"
	outPut += strings.Repeat("*", numStarsPerLine)
	return outPut
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {

	newMessage := strings.ReplaceAll(oldMsg, "*", " ")
	newMessage = strings.TrimSpace(newMessage)
	return newMessage
}
