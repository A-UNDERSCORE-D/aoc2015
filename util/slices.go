package util

import (
	"fmt"
	"reflect"
)

func SliceContains(slice interface{}, value interface{}) bool {
	sType := reflect.TypeOf(slice)
	vType := reflect.TypeOf(value)
	if sType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("SliceContains passed non slice %T", slice))
	}

	if sType.Elem().Kind() != vType.Kind() {
		panic(fmt.Sprintf("SliceContains passed non-matching slice and value types %T, %T", slice, value))
	}

	rSlice := reflect.ValueOf(slice)
	for i := 0; i < rSlice.Len(); i++ {
		v := rSlice.Index(i)
		if reflect.DeepEqual(v.Interface(), value) {
			return true
		}
	}
	return false
}

func StringSliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func IntSliceIndexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func IntSliceContains(slice []int, value int) bool {
	return IntSliceIndexOf(slice, value) != -1
}

func MakeIntSlice(length, step int) []int {
	num := 0
	out := make([]int, length)
	for i := 0; i < length; i++ {
		out[i] = num
		num += step
	}
	return out
}

func FilterStrSlice(in []string, filterFunc func(string) bool) []string {
	out := make([]string, 0, len(in))
	for _, v := range in {
		if filterFunc(v) {
			out = append(out, v)
		}
	}
	return out
}

func CopyIntSlice(a []int) []int {
	out := make([]int, len(a))
	copy(out, a)
	return out
}
