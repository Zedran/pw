package main

import (
	"strconv"
	"strings"
	"testing"
)

/* Tests GetLength function asserting it properly handles the length argument both for incorrect and
 * correct cases. Erratic input handling fails the test.
 */
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

/* Tests the alphanumeric generation, taking sample of generated chars and searching for those
 * on the extreme points of range. If they are not found, a warning is displayed. This test does not
 * trigger a fail due to its random input.
 */
func TestAlphanumeric(t *testing.T) {
	var sample []byte

	for i := 0; i < 10; i++ {
		chunk, err := randomStream(MIN_ALPHANUM_CODE, MAX_ALPHANUM_CODE, MAX_PW_LENGTH)
		if err != nil {
			t.Fatal(err)
		}
		sample = append(sample, chunk...)
	}

	strSample := string(sample)
	
	if !strings.Contains(strSample, string(MIN_ALPHANUM_CODE)) {
		t.Log("Check ASCII range. First code not found")
	}

	if !strings.Contains(strSample, string(MAX_ALPHANUM_CODE)) {
		t.Log("Check ASCII range. Last code not found")
	}
}

/* Tests the numeric generation, taking sample of generated numbers and searching for 0 and 9. 
 * If they are not found, a warning is displayed. This test does not trigger a fail due its random input.
 */
func TestNumeric(t *testing.T) {
	var sample []byte

	for i := 0; i < 10; i++ {
		chunk, err := randomStream(ASCII_ZERO, ASCII_NINE, MAX_PW_LENGTH)
		if err != nil {
			t.Fatal(err)
		}
		sample = append(sample, chunk...)
	}

	strSample := string(sample)

	if !strings.Contains(strSample, string(ASCII_ZERO)) {
		t.Log("Check numeric range. First code not found")
	}

	if !strings.Contains(strSample, string(ASCII_NINE)) {
		t.Log("Check numeric range. Last code not found")
	}
}
