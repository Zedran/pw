package main

import (
	"errors"
	"fmt"
	"log"
	"crypto/rand"
	"os"
	"strconv"
	"strings"
)

const (
	// The code of 0 for numeric function
	ASCII_ZERO        = 48

	// The next three constants are required for alphanumeric function

	// From '!'
	MIN_ALPHANUM_CODE = 33

	// Up to '}'
	MAX_ALPHANUM_CODE = 126

	// End range for rand.Intn function
	ALPHANUM_RANGE    = MAX_ALPHANUM_CODE - MIN_ALPHANUM_CODE

	// Maximum permitted password length
	MAX_PW_LENGTH     = 4096

	// When the error occurs, this string is printed together with error information
	ERR_INFO          = "\n%s\n\nAvailable modes: 'a' (alphanumeric password), 'n' (numeric password).\nIssue command with 'pw <mode> <length>'."
)

var (
	errUnknownArg         = errors.New("unknown argument encountered")
	errInvalidArgCount    = errors.New("too few or too many arguments passed on startup")

	errInvalidLengthArg   = errors.New("length must be an integer in range (0; 4096>")
)

// Generates alphanumeric password in range <33; 125) of ASCII codes
func alphanumeric(pwLength int) (string, error) {
	pw := make([]byte, pwLength)

	if _, err := rand.Read(pw); err != nil {
		return "", err
	}

	for i := range pw {
		pw[i] = MIN_ALPHANUM_CODE + pw[i] % ALPHANUM_RANGE
	}

	return string(pw), nil
}

// Generates numeric password in range <0; 9>
func numeric(pwLength int) (string, error) {
	pw := make([]byte, pwLength)

	if _, err := rand.Read(pw); err != nil {
		return "", err
	}

	for i := range pw {
		pw[i] = ASCII_ZERO + pw[i] % 10
	}

	return string(pw), nil
}

/* Gets password length from the argument passed and ensures it is non-negative integer lesser
 than predefined max length */
func getLength(lengthArg string) (int, error) {
	length, err := strconv.Atoi(lengthArg)
	
	if length <= 0 || length > MAX_PW_LENGTH || err != nil {
		return 0, errInvalidLengthArg
	}

	return length, nil
}

func main() {
	log.SetFlags(0)

	if len(os.Args) != 3 {
		log.Fatalf(ERR_INFO, errInvalidArgCount.Error())
	}

	length, err := getLength(os.Args[2])
	if err != nil {
		log.Fatalf(ERR_INFO, err.Error())
	}

	var pw string

	switch strings.ToUpper(os.Args[1]) {
	case "N":
		pw, err = numeric(length)
	case "A":
		pw, err = alphanumeric(length)
	default:
		log.Fatalf(ERR_INFO, errUnknownArg)
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nYour password: ", pw)
}
