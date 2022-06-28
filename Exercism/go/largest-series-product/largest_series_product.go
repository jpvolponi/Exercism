package lsproduct

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ConvInt(s string) []int64 {
	st_split_teste := strings.Split(s, "")
	var int_split []int64
	for _, v := range st_split_teste {
		if f, err := strconv.ParseInt(v, 10, 64); err == nil {
			int_split = append(int_split, f)

		} else {
			fmt.Println(err)
		}

	}
	return int_split
}
func LargestSeriesProduct(digits string, span int) (int64, error) {
	//var spanFloat float64 = float64(span)
	var digitsInt []int64 = ConvInt(digits)
	var biggest int64 = 1
	var aux int64 = 1

	if span <= 0 {
		return biggest, errors.New("Digits and span must be positive.")
	} else {
		for i := 0; i < len(digitsInt)-1; i++ {
			for j := 0; j < span-1; j++ {
				aux = digitsInt[i] * digitsInt[i+1]
				if biggest < aux {
					biggest = aux
				}
				fmt.Println("DEBUG aux:", aux)
				fmt.Println("DEBUG biggest:", biggest)
			}

		}
	}

	return biggest, nil
}
