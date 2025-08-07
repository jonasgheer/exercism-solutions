package reverse

func Reverse(input string) (reversed string) {
	for _, c := range input {
		reversed = string(c) + reversed
	}
	return reversed
}
