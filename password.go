package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"strings"
)

// compileCharset compiles a character set for password generation
// from the provided specifiers:
//   - A : upper case
//   - a : lower case
//   - n : numbers
//   - s : symbols
//
// The exclude variable contains characters that are excluded
// from the compiled charset.
//
// Returns an error if:
//   - set is an empty string
//   - set contains more than one specifier of the same kind
//   - set contains an undefined specifier
//   - exclude contains all characters from the selected charset
func compileCharset(set, exclude string) (string, error) {
	const (
		upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		lower   = "abcdefghijklmnopqrstuvwxyz"
		numbers = "0123456789"
		symbols = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	)

	if len(set) == 0 {
		return "", errors.New("empty set of specifiers")
	}

	var charset strings.Builder

	for _, c := range set {
		if strings.Count(set, string(c)) > 1 {
			return "", errors.New("duplicate specifier detected")
		}

		switch c {
		case 'A':
			charset.WriteString(upper)
		case 'a':
			charset.WriteString(lower)
		case 'n':
			charset.WriteString(numbers)
		case 's':
			charset.WriteString(symbols)
		default:
			return "", errors.New("undefined charset specifier")
		}
	}

	s := charset.String()
	for _, c := range exclude {
		s = strings.ReplaceAll(s, string(c), "")
	}

	if len(s) == 0 {
		return "", errors.New("empty charset due to exclusions")
	}

	return s, nil
}

// password generates a character password of the specified length.
// The charset used to generate a password is determined by byte values in set
// and the exclude variable, which contains the characters that are to be
// removed from the compiled charset.
func password(length int, set, exclude string) (string, error) {
	charset, err := compileCharset(set, exclude)
	if err != nil {
		return "", fmt.Errorf("charset compilation error: %w", err)
	}

	bytes, err := randBytes(length)
	if err != nil {
		return "", fmt.Errorf("random generator error: %w", err)
	}

	var (
		csLen = byte(len(charset))
		pw    strings.Builder
	)

	for _, b := range bytes {
		pw.WriteByte(charset[b%csLen])
	}

	return pw.String(), nil
}

// Returns a slice of random bytes of the specified length.
func randBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	return bytes, err
}
