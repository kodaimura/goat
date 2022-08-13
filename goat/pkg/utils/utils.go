package utils

import (
	"strconv"
)

func AtoiSlice(sl []string) ([]int, error) {
	isl := make([]int, len(sl))
	for i, v := range sl {
		x , err := strconv.Atoi(v)
		isl[i] = x
		if err != nil {
			return []int{}, err
		}
	}
	return isl, nil
}


func ItoaSlice(sl []int) []string {
	asl := make([]string, len(sl))
	for i, v := range sl {
		asl[i] = strconv.Itoa(v)
	}

	return asl
} 