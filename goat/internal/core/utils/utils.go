package utils

import (
	"math/rand"
	"time"
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


func RandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}