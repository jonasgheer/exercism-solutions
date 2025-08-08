package hamming

import "errors"

// Distance returns the hamming distance between a and b
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("a and b are of unequal length")
	}
	numDiffs := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			numDiffs++
		}
	}
	return numDiffs, nil
}
