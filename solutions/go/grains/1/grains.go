package grains

import (
	"fmt"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, fmt.Errorf("invalid square %d", number)
	}
	return uint64(math.Pow(2, float64(number)-1)), nil
}

func Total() uint64 {
	var total uint64 = 0
	for i := 1; i < 65; i++ {
		if res, err := Square(i); err == nil {
			total += res
		}
	}
	return total
}
