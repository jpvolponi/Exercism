package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var hammingDistance int

	if len(a) != len(b) {
		return hammingDistance, errors.New("Diferent DNA's strand length.")
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hammingDistance++
			}
		}
		return hammingDistance, nil
	}
}
