package allyourbase

import (
	"errors"
	"math"
)

func divmod(num int, den int) (int, int) {
	return num / den, num % den
}

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	switch {
	case inputBase < 2:
		return nil, errors.New("input base must be >= 2")
	case outputBase < 2:
		return nil, errors.New("output base must be >= 2")
	}
	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	sumBase10 := 0
	for i, j := 0, len(inputDigits)-1; i < len(inputDigits); i, j = i+1, j-1 {
		sumBase10 += int(float64(inputDigits[j]) * math.Pow(float64(inputBase), float64(i)))
	}

	quotient, remainder := divmod(sumBase10, outputBase)
	output := []int{remainder}
	for quotient != 0 {
		quotient, remainder = divmod(quotient, outputBase)
		output = append(output, remainder)
	}
	for i, j := 0, len(output)-1; i < len(output)/2; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}

	return output, nil
}
