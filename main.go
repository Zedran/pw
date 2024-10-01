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
		sep      = flag.String("s", DEFAULT_SEP, "word separator for diceware mode")
		wordList = flag.String("f", DEFAULT_WL, "file with word list for diceware mode")
	)

	flag.Parse()

	if !ValidateLength(*pwLen) {
		log.Fatalf("length must be an integer in range (0; %d>\n", MAX_PW_LENGTH)
	}

	var p string
	switch *mode {
	case "a":
		p = Alphanumeric(*pwLen)
	case "d":
		p = GeneratePhrases(*pwLen, *wordList, *sep)
	case "n":
		p = Numeric(*pwLen)
	default:
		log.Fatal("invalid mode argument\n")
	}
	fmt.Println(p)
}
