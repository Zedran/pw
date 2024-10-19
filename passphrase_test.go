package main

import (
	"strings"
	"testing"

	"github.com/Zedran/pw/internal/tests"
)

// Tests Generate Phrases function for the correct number of words generated.
// Incorrect words count fails the test.
func TestGeneratePhrases(t *testing.T) {
	const sep = "_sep_"

	cases := []int{1, 4, 75, MAX_PW_LENGTH}

	for i := range cases {
		output, err := passphrase(cases[i], DEFAULT_WL, sep)
		if err != nil {
			t.Fatalf("failed to generate passphrase: %v", err)
		}

		if len(strings.Split(output, sep)) != cases[i] {
			t.Errorf("Failed for case '%d'. Output: '%s'", cases[i], output)
		}
	}
}

// Tests GetRandomNumbers function. Takes a sample of 300000 random numbers
// and checks whether a correct range of numbers is represented. The test
// displays warnings if the extreme values are not present or if the number
// falls outside the range <0, wordPool). Due to the nature of the random
// output, not meeting a criteria does not fail the test.
func TestGetRandomNumbers(t *testing.T) {
	const (
		wordPool   = 7776 // The length of a list for 5 dices
		sampleSize = 300000
	)

	sample, err := GetRandomNumbers(sampleSize, int(wordPool))
	if err != nil {
		t.Fatalf("failed to generate random numbers: %v", err)
	}

	var (
		zeroPresent  = false
		maxPresent   = false
		correctRange = true
	)

	for i := range sample {
		switch sample[i] {
		case 0:
			zeroPresent = true
		case wordPool - 1:
			maxPresent = true
		default:
			if sample[i] < 0 || sample[i] >= wordPool {
				correctRange = false
			}
		}
	}

	if !(zeroPresent && maxPresent && correctRange) {
		t.Logf(
			"Range criteria not met for random sample: Min: %t, Max: %t, Over or under: %t",
			zeroPresent,
			maxPresent,
			correctRange,
		)
	}
}

// Tests LoadWordList function. Inquires whether the loaded slice has any empty
// lines or are there new lines at the end of any of the elements. Incorrect
// word list format fails the test.
func TestLoadWordList(t *testing.T) {
	var (
		emptyLinePresent = false
		newLinesPresent  = false
	)

	wordList, err := readWordList(DEFAULT_WL)
	if err != nil {
		t.Fatalf("Failed to load word list: %v", err)
	}

	for i := range wordList {
		if len(strings.TrimSpace(wordList[i])) == 0 {
			emptyLinePresent = true
		} else if strings.Contains(wordList[i], "\n") {
			newLinesPresent = true
		}
	}

	if emptyLinePresent || newLinesPresent {
		t.Errorf(
			"Word list did not load correctly. Empty lines present: %t, Trailing new lines: %t",
			emptyLinePresent,
			newLinesPresent,
		)
	}
}

func TestPassphraseBias(t *testing.T) {
	const (
		P_LEN       = 10
		SAMPLE_SIZE = 10_000
		CRIT_VALUE  = tests.CRIT_PASSPHRASE

		// chisqr <|> CRIT_VALUE [biased]
		FMT = "%.3f (%s%.3f) %s\n"
	)

	wl, err := readWordList(DEFAULT_WL)
	if err != nil {
		t.Fatalf("failed to load word list: %v", err)
	}

	occurences := make([]float64, len(wl))

	for range SAMPLE_SIZE {
		pp, err := passphrase(P_LEN, DEFAULT_WL, DEFAULT_SEP)
		if err != nil {
			t.Fatalf("failed to generate passphrase")
		}

		s := strings.Split(pp, DEFAULT_SEP)

		for i, w := range wl {
			occurences[i] += tests.Count(s, w)
		}
	}

	if chisqr, biased := tests.SampleBiased(occurences, CRIT_VALUE); biased {
		t.Fatalf(FMT, chisqr, ">", CRIT_VALUE, "-- biased")
	} else {
		t.Logf(FMT, chisqr, "<", CRIT_VALUE, "")
	}
}
