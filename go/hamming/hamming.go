package hamming

import "errors"

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) != len(b) {
		return 0, errors.New("the string must be the same length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil

}
