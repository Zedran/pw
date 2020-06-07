package main

import (
	"strconv"
	"strings"
	"testing"
)

/* Tests GetLength function asserting it properly handles the length argument both for incorrect and
 correct cases. Erratic input handling fails the test */
func TestGetLength(t *testing.T) {
	errorCases := []string{
		"0",                      // zero
		"-11",                    // negative
		"17.4",                   // float 
		"4097",                   // out of predefined bounds
		"-92233720368547758070",  // negative out of the int bounds
		"92233720368547758070",   // positive out of the int bounds
		"a",                      // nonconvertible char
		"ala",                    // more nonconvertible chars
		" ",                      // whitespace
		"",                       // empty string
	}

	for _, testCase := range errorCases {
		if _, err := getLength(testCase); err == nil {
			t.Errorf("Failed for %s\n", testCase)
		}
	}
	
	correctCases := map[string]int{
	    "1": 1,                                      // minimum
		"15": 15,                                    // some reasonable length
		strconv.Itoa(MAX_PW_LENGTH): MAX_PW_LENGTH,  // maximum
	}

	for k, v := range correctCases {
		x, err := getLength(k)
		if err != nil || x != v {
			t.Errorf("Failed for %s - expected: %d, outcome: %d\n", k, v, x)
		}
	}
}

/* Test the alphanumeric function, taking sample of generated chars and searching for those
 on the extreme points of range. If they are not found, a warning is displayed. This test does not
 trigger fail due to its random input */
func TestAlphanumeric(t *testing.T) {
	var bigString string

	for i := 0; i < 10; i++ {
		bigString += alphanumeric(MAX_PW_LENGTH)
	}

	if !strings.Contains(bigString, string(MIN_ALPHANUM_CODE)) {
		t.Log("Check ASCII range. First code not found")
	}

	if !strings.Contains(bigString, string(MAX_ALPHANUM_CODE - 1)) {
		t.Log("Check ASCII range. Last code not found")
	}
}

/* Test the numeric function, taking sample of generated numbers and searching for numbers
 on the extreme points of range. If they are not found, a warning is displayed. This test does not
 trigger fail due its random input */
func TestNumeric(t *testing.T) {
	var bigString string

	for i := 0; i < 10; i++ {
		bigString += numeric(MAX_PW_LENGTH)
	}

	if !strings.Contains(bigString, string(ASCII_ZERO)) {
		t.Log("Check numeric range. First code not found")
	}

	if !strings.Contains(bigString, string(ASCII_ZERO + 9)) {
		t.Log("Check numeric range. Last code not found")
	}
}
