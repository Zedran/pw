package main

import (
	"flag"
	"fmt"
	"log"
)

const (
	// Alphanumeric bounds
	MIN_ALPHANUM_CODE = 33  // From '!'
	MAX_ALPHANUM_CODE = 125 // Up to '}'

	// Numeric bounds
	ASCII_ZERO = 48
	ASCII_NINE = ASCII_ZERO + 9

	// Maximum permitted password length
	MAX_PW_LENGTH = 4096

	// Default password length
	DEFAULT_LENGTH = 6
)

func main() {
	log.SetFlags(0)

	mode := flag.String("m", "a", "mode:\n    a - alphanumeric\n    d - diceware\n    n - numeric")
	pwLen := flag.Int("l", DEFAULT_LENGTH, fmt.Sprintf("password length, max %d characters", MAX_PW_LENGTH))
	wordList := flag.String("f", DEFAULT_WL, "file with word list for diceware mode")

	flag.Parse()

	if !ValidateLength(*pwLen) {
		log.Fatalf("length must be an integer in range (0; %d>\n", MAX_PW_LENGTH)
	}

	switch *mode {
	case "a": // Alphanumeric
		fmt.Println(Alphanumeric(*pwLen))
	case "d": // Diceware
		fmt.Println(GeneratePhrases(*pwLen, *wordList))
	case "n": // Numeric
		fmt.Println(Numeric(*pwLen))
	default:
		log.Fatal("invalid mode argument\n")
	}
}
