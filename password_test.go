package main

import (
	"strings"
	"testing"

	"github.com/Zedran/pw/internal/tests"
)

func TestPasswordBias(t *testing.T) {
	const (
		P_LEN       = 20
		SET         = "Aans"
		SAMPLE_SIZE = 10_000
		CRIT_VALUE  = tests.CRIT_PASSWORD

		// chisqr <|> CRIT_VALUE [biased]
		FMT = "%.3f (%s%.3f) %s\n"
	)

	charset, err := compileCharset(SET, "")
	if err != nil {
		t.Fatalf("charset compilation error: %v", err)
	}

	occurences := make([]float64, len(charset))

	for range SAMPLE_SIZE {
		pw, err := password(P_LEN, SET, "")
		if err != nil {
			t.Fatalf("password function returned an error: %v", err)
		}

		for i, b := range charset {
			occurences[i] += float64(strings.Count(pw, string(b)))
		}
	}

	if chisqr, biased := tests.SampleBiased(occurences, CRIT_VALUE); biased {
		t.Fatalf(FMT, chisqr, ">", CRIT_VALUE, "-- biased")
	} else {
		t.Logf(FMT, chisqr, "<", CRIT_VALUE, "")
	}
}

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
