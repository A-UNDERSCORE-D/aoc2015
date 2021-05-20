package util

import (
	"math/bits"
)

func BitExplode(num uint) []int {
	bLen := bits.Len(num)
	out := make([]int, 0, bLen)
	for i := 1; i <= bLen; i++ {
		out = append(out, int(num&uint(1<<(bLen-i)))>>(bLen-i))
	}
	return out
}

func DecimalExplode(num int) []int {
	out := make([]int, 0, (num/10)+1)
	for num > 9 {
		out = append(out, num%10)
		num = num / 10
	}
	out = append(out, num)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}
