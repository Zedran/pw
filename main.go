package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)

	var (
		entropy  = flag.Bool("E", false, "Calculate entropy for password with selected parameters and exit")
		exclude  = flag.String("e", "", "Exclude characters in password mode")
		include  = flag.String("i", "Aans", "Include character groups in charset for the password generator:\n\tA - upper case\n\ta - lower case\n\tn - numbers\n\ts - symbols\n")
		mode     = flag.String("m", "c", "Generation mode:\n\tc - password\t(characters)\n\tw - passphrase\t(words)\n\n")
		noLF     = flag.Bool("n", false, "Do not print an LF character at the end")
		pwLen    = flag.Int("l", DEFAULT_LENGTH, fmt.Sprintf("Number of generated elements: <1; %d> characters or words", MAX_PW_LENGTH))
		sep      = flag.String("s", DEFAULT_SEP, "Word separator for passphrase mode")
		wordList = flag.String("f", DEFAULT_WL, "Word list file for passphrase mode")
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
		if !*entropy {
			p, err = password(*pwLen, *include, *exclude)
		} else {
			n, err := entropyC(*pwLen, *include, *exclude)
			if err != nil {
				log.Fatal(err)
			}
			p = fmt.Sprintf("%.2f", n)
		}
	case "w":
		if !*entropy {
			p, err = passphrase(*pwLen, *wordList, *sep)
		} else {
			n, err := entropyW(*pwLen, *wordList)
			if err != nil {
				log.Fatal(err)
			}
			p = fmt.Sprintf("%.2f", n)
		}
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
