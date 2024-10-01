package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)

	var (
		mode     = flag.String("m", "a", "mode:\n    a - alphanumeric\n    d - diceware\n    n - numeric")
		pwLen    = flag.Int("l", DEFAULT_LENGTH, fmt.Sprintf("password length, max %d characters", MAX_PW_LENGTH))
		wordList = flag.String("f", DEFAULT_WL, "file with word list for diceware mode")
	)

	flag.Parse()

	if !ValidateLength(*pwLen) {
		log.Fatalf("length must be an integer in range (0; %d>\n", MAX_PW_LENGTH)
	}

	switch *mode {
	case "a":
		fmt.Println(Alphanumeric(*pwLen))
	case "d":
		fmt.Println(GeneratePhrases(*pwLen, *wordList))
	case "n":
		fmt.Println(Numeric(*pwLen))
	default:
		log.Fatal("invalid mode argument\n")
	}
}
