package iban

func mod97(number string) int {
	var m = 0
	for _, c := range number {
		if c < '0' || c > '9' {
			return -1
		}
		m = (((m * 10) + (int(c) - '0')) % 97)
	}
	return m
}
