package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)

	var (
		include  = flag.String("i", "Aans", "Charset for character password:\n    A - upper case\n    a - lower case\n    n - numbers\n    s - symbols\n")
		mode     = flag.String("m", "c", "Generation mode:\n    c - character password\n    w - phrase password\n\n")
		noLF     = flag.Bool("n", false, "Do not print an LF character at the end")
		pwLen    = flag.Int("l", DEFAULT_LENGTH, fmt.Sprintf("Password length, <1; %d> characters", MAX_PW_LENGTH))
		sep      = flag.String("s", DEFAULT_SEP, "Word separator for diceware mode")
		wordList = flag.String("f", DEFAULT_WL, "Word list file for diceware mode")
	)

	flag.Parse()

	if !ValidateLength(*pwLen) {
		log.Fatalf("length must be an integer in range <1; %d>\n", MAX_PW_LENGTH)
	}

	var (
		p   string
		err error
	)

	switch *mode {
	case "c":
		p, err = password(*pwLen, *include)
	case "w":
		p = GeneratePhrases(*pwLen, *wordList, *sep)
	default:
		log.Fatal("invalid mode argument\n")
	}

	if err != nil {
		log.Fatal(err)
	}

	if *noLF {
		fmt.Print(p)
	} else {
		fmt.Println(p)
	}
}
