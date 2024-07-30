package helpers

import "strconv"

func IsNumber(num string) (int, bool) {
	n, err := strconv.Atoi(num)
	return n, err == nil
}

func NumberIsInRange(n int, min, max int) bool {
	return n >= min && n <= max
}
