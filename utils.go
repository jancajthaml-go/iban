package iban

func mod97(number string) int {
	var (
		d uint
		i int
		x uint
		l int = len(number)
	)

scan:
	d = uint(number[i]) - 48
	if d > 9 {
		return -1
	}
	x = (((x * 10) + d) % 97)
	i++
	if i != l {
		goto scan
	}
	return int(x)
}
