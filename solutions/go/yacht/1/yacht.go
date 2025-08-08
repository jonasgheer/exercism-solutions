package yacht

const numDice = 5

func Score(dice []int, category string) int {
	diceCount := make(map[int]int)
	for _, d := range dice {
		diceCount[d]++
	}

	switch category {
	case "ones":
		return diceCount[1]
	case "twos":
		return diceCount[2] * 2
	case "threes":
		return diceCount[3] * 3
	case "fours":
		return diceCount[4] * 4
	case "fives":
		return diceCount[5] * 5
	case "sixes":
		return diceCount[6] * 6
	case "full house":
		sum := 0
		for d, c := range diceCount {
			if c != 2 && c != 3 {
				return 0
			}
			sum += d * c
		}
		return sum
	case "four of a kind":
		for d, c := range diceCount {
			if c >= 4 {
				return d * 4
			}
		}
	case "little straight":
		for i := 1; i <= 5; i++ {
			if _, exist := diceCount[i]; !exist {
				return 0
			}
		}
		return 30
	case "big straight":
		for i := 2; i <= 6; i++ {
			if _, exist := diceCount[i]; !exist {
				return 0
			}
		}
		return 30
	case "choice":
		sum := 0
		for _, d := range dice {
			sum += d
		}
		return sum
	case "yacht":
		if len(diceCount) == 1 {
			return 50
		}
	}
	return 0
}
