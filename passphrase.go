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

// Builds a phrase of the specified length from word list file
// at the specified path. Words are separated by sep.
func passphrase(length int, wordListPath, sep string) string {
	var phrases = make([]string, length)

	wordList, err := readWordList(wordListPath)
	if err != nil {
		log.Fatalf("Failed to load %s: %v", wordListPath, err)
	}

	numbers := GetRandomNumbers(length, len(wordList))

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

// Returns a slice of words for passphrase generator, reading it from
// the embedded file system or OS path.
func readWordList(path string) ([]string, error) {
	var load func(string) ([]byte, error)

	if path == DEFAULT_WL {
		load = efs.ReadFile
	} else {
		load = os.ReadFile
	}

	stream, err := load(path)

	return strings.Split(string(stream), "\n"), err
}
