package tests

func Count(slice []string, target string) float64 {
	var count float64 = 0

	for _, e := range slice {
		if e == target {
			count++
		}
	}
	return count
}
