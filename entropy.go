package main

import (
	"fmt"
	"math"
)

// entropyC calculates entropy of a password (individual characters)
// of given length and charset.
func entropyC(length int, set, exclude string) (float64, error) {
	charset, err := compileCharset(set, exclude)
	if err != nil {
		return -1, fmt.Errorf("charset compilation error: %w", err)
	}
	return getEntropy(len(charset), length), nil
}

// entropyW calculates entropy of a passphrase (individual words),
// of given length and word pool.
func entropyW(length int, wordListPath string) (float64, error) {
	wordList, err := readWordList(wordListPath)
	if err != nil {
		return -1, err
	}
	return getEntropy(len(wordList), length), nil
}

// getEntropy calculates entropy bits given the count of unique elements
// in a set and the length of a subset.
func getEntropy(setLen, subsetLen int) float64 {
	return math.Log2(float64(setLen)) * float64(subsetLen)
}
