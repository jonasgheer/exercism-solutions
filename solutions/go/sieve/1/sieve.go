package sieve

func Sieve(limit int) []int {
	sequence := make([]bool, limit+1)
	primes := []int{}

	for n := 2; n <= limit; n++ {
		marked := sequence[n]
		if !marked {
			primes = append(primes, n)
			for i := 1; ; i++ {
				multiple := n * i
				if multiple > limit {
					break
				}
				sequence[multiple] = true
			}
		}
	}
	return primes
}
