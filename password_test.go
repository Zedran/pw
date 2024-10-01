package main

import (
	"strings"
	"testing"
)

func TestCompileCharset(t *testing.T) {
	errCases := []string{
		"",    // Empty set
		"AA",  // Duplicate specifier
		"aAd", // Undefined specifier
	}

	for _, c := range errCases {
		if out, err := compileCharset(c, ""); err == nil {
			t.Fatalf("error not returned, output: %s", out)
		}
	}

	set, err := compileCharset("n", "1")
	if err != nil {
		t.Fatalf("error returned for set 'n', exclusions '1': %v", err)
	}

	if strings.Count(set, "1") > 0 {
		t.Fatalf("'1' not excluded from charset: %s", set)
	}

	if _, err = compileCharset("n", "0123456789"); err == nil {
		t.Fatalf("complete charset exclusion did not cause an error")
	}
}
