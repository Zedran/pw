package main

import (
	"strings"
	"testing"
)

/* Tests Generate Phrases function for the correct number of words generated. 
 * Incorrect words count fails the test.
 */
func TestGeneratePhrases(t *testing.T) {
	cases := []int{1, 4, 75, MAX_PW_LENGTH}

	var output string
	for i := range output {
		output = string(GeneratePhrases(cases[i], DEFAULT_WL))

		if len(strings.Split(output, OUTPUT_SEP)) != cases[i] {
			t.Errorf("Failed for case '%d'. Output: '%s'", cases[i], output)
		}
	}
}

/* Tests GetRandomNumbers function. Takes a sample of 300000 random numbers and checks whether a correct range 
 * of numbers is represented. The test displays warnings if the extreme values are not present
 * or if the number falls outside the range <0, wordPool). Due to the nature of the random output, not meeting
 * a criteria does not fail the test.
 */
func TestGetRandomNumbers(t *testing.T) {
	var (
		wordPool   int64 = 7776    // the length of a list for 5 dices
		sampleSize       = 300000

		sample           = GetRandomNumbers(sampleSize, int(wordPool))
		
		zeroPresent      = false
		maxPresent       = false
		correctRange     = true
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

/* Tests LoadWordList function. Inquires whether the loaded slice has any empty lines or are there 
 * new lines at the end of any of the elements. Incorrect word list format fails the test.
 */
func TestLoadWordList(t *testing.T) {
	var (
		wordList         = LoadWordList(DEFAULT_WL)
		emptyLinePresent = false
		newLinesPresent  = false
	)

	for i := range wordList {
		if len(strings.TrimSpace(wordList[i])) == 0 {
			emptyLinePresent = true
		} else if strings.Contains(wordList[i], WL_SEP) {
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
