package util

import (
	"strconv"
	"strings"
)

// ReplaceAllSlice uses strings.ReplaceAll on each index of the given slice. The slice IS modified, but also returned
func ReplaceAllSlice(in []string, old, new string) []string {
	for i, v := range in {
		in[i] = strings.ReplaceAll(v, old, new)
	}

	return in
}

func ContainsAny(in string, substrs ...string) bool {
	for _, substr := range substrs {
		if strings.Contains(in, substr) {
			return true
		}
	}

	return false
}

func MustAtoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return res
}
