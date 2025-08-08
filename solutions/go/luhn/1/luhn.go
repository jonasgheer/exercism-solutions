package luhn

import (
    "strings"
    "strconv"
)

func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")
    if len(id) <= 1 {
        return false
    }
    idSlice := strings.Split(id, "")
    idNum := make([]int, len(idSlice))
    for _, c := range idSlice {
        n, err := strconv.Atoi(c)
        if err != nil {
            return false
        }
        idNum = append(idNum, n)
    }
    for i := len(idNum) - 2; i >= 0; i -= 2 {
        double := idNum[i] * 2
        if double > 9 {
            double -= 9
        }
        idNum[i] = double
    }
	sum := 0
    for _, n := range idNum {
        sum += n
    }
	return sum % 10 == 0
}
