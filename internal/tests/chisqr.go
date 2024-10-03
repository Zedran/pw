package tests

import "math"

const (
	// 7776 elements (df = 7775)
	CRIT_PASSPHRASE float64 = 7935.23115 // 0.90

	// 97 elements (df = 96)
	CRIT_PASSWORD float64 = 114.13071 // 0.90
)

func SampleBiased(occurences []float64, crit float64) (float64, bool) {
	chisqr := chisqr(occurences)
	return chisqr, chisqr > crit
}

func chisqr(occurences []float64) float64 {
	var (
		chisqr     float64 = 0
		sampleSize float64 = 0
	)

	for _, n := range occurences {
		sampleSize += n
	}

	expected := sampleSize / float64(len(occurences))

	for _, occ := range occurences {
		chisqr += math.Pow(occ-expected, 2)
	}

	chisqr /= expected

	return chisqr
}
