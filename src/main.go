package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	// Alphanumeric bounds
	MIN_ALPHANUM_CODE = 33      // From '!'
	MAX_ALPHANUM_CODE = 125    // Up to '}'

	// Numeric bounds
	ASCII_ZERO        = 48
	ASCII_NINE        = ASCII_ZERO + 9

	// Maximum permitted password length
	MAX_PW_LENGTH     = 4096

	// A secret random number for the program to tell if the user has provided password length
	DEFAULT_LENGTH    = 6
)

func displayErrorMessage(message string) {
	fmt.Println(message)
	flag.PrintDefaults()
	os.Exit(1)
}

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

/* Returns true if length argument is within set bounds. */
func ValidateLength(length int) bool {
	return length > 0 && length <= MAX_PW_LENGTH
}

func main() {
	log.SetFlags(0)

	mode  := flag.String("m", "a", "mode:\n    a - alphanumeric\n    n - numeric")
	pwLen := flag.Int("l", DEFAULT_LENGTH, fmt.Sprintf("password length, max %d characters", MAX_PW_LENGTH))

	flag.Parse()

	if ValidateLength(*pwLen) == false {
		displayErrorMessage(fmt.Sprintf("length must be an integer in range (0; %d>\n", MAX_PW_LENGTH))
	}

	var (
		pw  []byte
		err error
	)

	switch *mode {
	case "a":  // alphanumeric
		pw, err = randomStream(MIN_ALPHANUM_CODE, MAX_ALPHANUM_CODE, *pwLen)
	case "n":  // numeric
		pw, err = randomStream(ASCII_ZERO, ASCII_NINE, *pwLen)
	default:
		displayErrorMessage("invalid mode argument\n")
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(pw))
}
