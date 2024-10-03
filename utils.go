package main

const (
	// Default password length
	DEFAULT_LENGTH = 6

	// Maximum permitted password length
	MAX_PW_LENGTH = 128
)

// Returns true if length argument is within set bounds.
func ValidateLength(length int) bool {
	return length > 0 && length <= MAX_PW_LENGTH
}
