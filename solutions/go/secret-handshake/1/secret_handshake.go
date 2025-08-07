package secret

func hasBit(n uint, pos int) bool {
	val := n & (1 << pos)
	return val > 0
}

var codeWords = map[int]string{
	0: "wink",
	1: "double blink",
	2: "close your eyes",
	3: "jump",
}

func Handshake(code uint) (handshake []string) {
	for i := 0; i < 4; i++ {
		if hasBit(code, i) {
			handshake = append(handshake, codeWords[i])
		}
	}
	if hasBit(code, 4) {
		for i := 0; i < len(handshake)/2; i++ {
			j := len(handshake) - i - 1
			handshake[i], handshake[j] = handshake[j], handshake[i]
		}
	}
	return handshake
}
