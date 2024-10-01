package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)

	var (
		mode     = flag.String("m", "a", "Generation mode:\n    a - alphanumeric\n    d - diceware\n    n - numeric\n")
		pwLen    = flag.Int("l", DEFAULT_LENGTH, fmt.Sprintf("Password length, <1; %d> characters", MAX_PW_LENGTH))
		sep      = flag.String("s", DEFAULT_SEP, "Word separator for diceware mode")
		wordList = flag.String("f", DEFAULT_WL, "Word list file for diceware mode")
	)

	flag.Parse()

	if !ValidateLength(*pwLen) {
		log.Fatalf("length must be an integer in range <1; %d>\n", MAX_PW_LENGTH)
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
