package main

import (
	"crypto/rand"
	"log"
)

const (
	// Alphanumeric bounds
	MIN_ALPHANUM_CODE = 33  // From '!'
	MAX_ALPHANUM_CODE = 125 // Up to '}'

	// Numeric bounds
	ASCII_ZERO = 48
	ASCII_NINE = ASCII_ZERO + 9
)

// Returns a random string of alphanumeric and sign characters.
func Alphanumeric(length int) string {
	return randomStream(MIN_ALPHANUM_CODE, MAX_ALPHANUM_CODE, length)
}

// Returns a random string of numeric characters.
func Numeric(length int) string {
	return randomStream(ASCII_ZERO, ASCII_NINE, length)
}

// Generates random stream of bytes and transforms them to fit
// within the specified range.
func randomStream(min, max, length int) string {
	stream := make([]byte, length)

	if _, err := rand.Read(stream); err != nil {
		log.Fatal(err)
	}

	for i := range stream {
		stream[i] = byte(min) + stream[i]%byte(max+1-min)
	}

	return string(stream)
}

// Returns true if length argument is within set bounds.
func ValidateLength(length int) bool {
	return length > 0 && length <= MAX_PW_LENGTH
}
