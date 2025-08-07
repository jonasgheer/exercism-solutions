package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	results := make(chan FreqMap, len(l))

	for _, line := range l {
		go func(line string) { results <- Frequency(line) }(line)
	}

	totalResult := make(FreqMap, 26)
	for i := 0; i < len(l); i++ {
		result := <-results
		for k, v := range result {
			totalResult[k] += v
		}
	}

	return totalResult
}
