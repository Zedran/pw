package main

import "testing"

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
