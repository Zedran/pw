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
	EXPECTED_ARGC     = 3

	// Alphanumeric bounds
	MIN_ALPHANUM_CODE = 33      // From '!'
	MAX_ALPHANUM_CODE = 125    // Up to '}'

	// Numeric bounds
	ASCII_ZERO        = 48
	ASCII_NINE        = ASCII_ZERO + 9

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

/* Generates random stream of bytes and transforms them to fit within the specified range. */
func randomStream(min, max, length int) ([]byte, error) {
	pw := make([]byte, length)

	if _, err := rand.Read(pw); err != nil {
		return nil, err
	}

	for i := range pw {
		pw[i] = byte(min) + pw[i] % byte(max + 1 - min)
	}

	return pw, nil
}

/* Gets password length from the string argument passed and ensures it is a non-negative integer lesser
 * than predefined max length.
 */
func getLength(lengthArg string) (int, error) {
	length, err := strconv.Atoi(lengthArg)
	
	if length <= 0 || length > MAX_PW_LENGTH || err != nil {
		return 0, errInvalidLengthArg
	}

	return length, nil
}

func main() {
	log.SetFlags(0)

	if len(os.Args) != EXPECTED_ARGC {
		log.Fatalf(ERR_INFO, errInvalidArgCount.Error())
	}

	length, err := getLength(os.Args[2])
	if err != nil {
		log.Fatalf(ERR_INFO, err.Error())
	}

	var pw []byte

	switch strings.ToUpper(os.Args[1]) {
	case "N":     // numeric
		pw, err = randomStream(ASCII_ZERO, ASCII_NINE, length)
	case "A":     // alphanumeric
		pw, err = randomStream(MIN_ALPHANUM_CODE, MAX_ALPHANUM_CODE, length)
	default:
		log.Fatalf(ERR_INFO, errUnknownArg)
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nYour password: ", string(pw))
}
