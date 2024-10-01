package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"os"
	"strings"
)

// File name of the default word list
const (
	DEFAULT_SEP string = " "
	DEFAULT_WL  string = "embed/eff"
)

// Generates a string containing random space-separated words for the user to
// join together in a password. Takes 3 arguments: a count of words,
// the name of the file containing a word list and word separator.
func GeneratePhrases(count int, wordListName, sep string) string {
	var phrases = make([]string, count)

	wordList, err := loadWordList(wordListName)
	if err != nil {
		log.Fatalf("Failed to load %s: %v", wordListName, err)
	}

	numbers := GetRandomNumbers(count, len(wordList))

	for i := range numbers {
		phrases[i] = wordList[numbers[i]]
	}

	return strings.Join(phrases, sep)
}

// Returns a slice of random integers range <0, wordPool).
func GetRandomNumbers(count int, wordPool int) []int64 {
	var (
		max  = big.NewInt(int64(wordPool))
		nums = make([]int64, count)
	)

	for i := range nums {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			log.Fatal(err)
		}

		nums[i] = n.Int64()
	}

	return nums
}

// Returns a slice of words for diceware generator, reading it from
// the embedded file system or OS path.
func loadWordList(path string) ([]string, error) {
	var load func(string) ([]byte, error)

	if path == DEFAULT_WL {
		load = efs.ReadFile
	} else {
		load = os.ReadFile
	}

	stream, err := load(path)

	return strings.Split(string(stream), "\n"), err
}
