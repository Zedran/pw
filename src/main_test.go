package main

import (
	"strings"
	"testing"
)

// Tests ValidateLength function asserting it properly checks the length
// argument both for incorrect and correct cases. Erratic input checking
// fails the test.
func TestValidateLength(t *testing.T) {
	errorCases := []int{
		0,    // Zero
		-11,  // Negative
		4097, // Out of predefined bounds by 1
	}

	for _, testCase := range errorCases {
		if ValidateLength(testCase) != false {
			t.Errorf("Failed for the incorrect case '%d'\n", testCase)
		}
	}

	correctCases := []int{
		1,             // Minimum
		15,            // Some reasonable length
		MAX_PW_LENGTH, // Maximum
	}

	for _, testCase := range correctCases {
		if ValidateLength(testCase) != true {
			t.Errorf("Failed for the correct case '%d'\n", testCase)
		}
	}
}

// Tests the alphanumeric generation, taking sample of generated chars
// and searching for those on the extreme points of range. If they are
// not found, a warning is displayed. This test does not trigger a fail
// due to its random input.
func TestAlphanumeric(t *testing.T) {
	var sample []byte

	for i := 0; i < 10; i++ {
		sample = append(sample, randomStream(MIN_ALPHANUM_CODE, MAX_ALPHANUM_CODE, MAX_PW_LENGTH)...)
	}

	strSample := string(sample)

	if !strings.Contains(strSample, string(byte(MIN_ALPHANUM_CODE))) {
		t.Log("Check ASCII range - first code not found")
	}

	if !strings.Contains(strSample, string(byte(MAX_ALPHANUM_CODE))) {
		t.Log("Check ASCII range - last code not found")
	}
}

// Tests the numeric generation, taking sample of generated numbers
// and searching for 0 and 9. If they are not found, a warning is displayed.
// This test does not trigger a fail due to its random input.
func TestNumeric(t *testing.T) {
	var sample []byte

	for i := 0; i < 10; i++ {
		sample = append(sample, randomStream(ASCII_ZERO, ASCII_NINE, MAX_PW_LENGTH)...)
	}

	strSample := string(sample)

	if !strings.Contains(strSample, string(byte(ASCII_ZERO))) {
		t.Log("Check numeric range. First code not found")
	}

	if !strings.Contains(strSample, string(byte(ASCII_NINE))) {
		t.Log("Check numeric range. Last code not found")
	}
}
