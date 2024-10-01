package main

import "testing"

func TestCompileCharset(t *testing.T) {
	errCases := []string{
		"",    // Empty set
		"AA",  // Duplicate specifier
		"aAd", // Undefined specifier
	}

	for _, c := range errCases {
		if out, err := compileCharset(c); err == nil {
			t.Fatalf("error not returned, output: %s", out)
		}
	}
}
