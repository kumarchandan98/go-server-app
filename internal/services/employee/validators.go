package employee

import "net/mail"

func ValidateEmail(email string) bool {
	if _, err := mail.ParseAddress(email); err != nil {
		return false
	}
	return true
}
func ValidateRole(role string) bool {
	if role != "developer" && role != "manager" {
		return false
	} else {
		return true
	}
}