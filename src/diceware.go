package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"os"
	"path"
	"strings"
)

const (
	// Word separator inside the word list file
	WL_SEP     string = "\n"

	// Word separator in the output string
	OUTPUT_SEP string = " "

	// Directory containing word list files
	WL_DIR     string = "res"

	// Extension of word list files
	WL_EXT     string = ".txt"

	// File name of the default word list
	DEFAULT_WL string = "eff"
)

/* Generates a string containing random space-separated words for the user to join together in a password. 
 * Takes 2 arguments: a count of words and the name of the file containing a word list.
 */
func GeneratePhrases(count int, wordListName string) string {
	var (
		wordList = LoadWordList(wordListName)
		numbers  = GetRandomNumbers(count, len(wordList))
		phrases  = make([]string, count)
	)

	for i := range numbers {
		phrases[i] = wordList[numbers[i]]
	}
	
	return strings.Join(phrases, OUTPUT_SEP)
}

/* Returns a slice of random integers range <0, wordPool). */
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

/* Loads a list of words from a word list file. */
func LoadWordList(name string) []string {
	stream, err := os.ReadFile(path.Join(WL_DIR, name + WL_EXT))
	if err != nil {
		log.Fatal(err)
	}

	var wordList []string
	for _, word := range strings.Split(string(stream), WL_SEP) {
		if len(strings.TrimSpace(word)) != 0 {
			wordList = append(wordList, word)
		}
	}

	return wordList
}
