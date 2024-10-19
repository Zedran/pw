package main

import (
	"math"
	"testing"
)

func TestGetEntropy(t *testing.T) {
	type testCase struct {
		charset    string
		exclusions string
		length     int
		expected   float64
	}

	cases := []testCase{
		{"A", "", 1, 4.7004397},
		{"Aa", "", 1, 5.7004397},
		{"Aan", "", 1, 5.9541963},
		{"Aans", "", 1, 6.5698556},
		{"Aans", "_", 1, 6.5545888},
		{"Aans", "c", 3, 19.6637665},
		{"Aans", "", 128, 840.9415178},
	}

	const tolerance float64 = 1e-6

	for i, c := range cases {
		charset, err := compileCharset(c.charset, c.exclusions)
		if err != nil {
			t.Fatalf("charset generation error for case %d: %v", i, err)
		}

		e := getEntropy(len(charset), c.length)

		if math.Abs(e-c.expected) > tolerance {
			t.Fatalf("Entropy value discrepancy for case %d: expected %f, got %f", i, c.expected, e)
		}
	}
}
