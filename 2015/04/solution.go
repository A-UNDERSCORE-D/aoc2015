package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

	"awesome-dragon.science/go/adventofcode2015/util"
)

func main() {
	input := util.ReadLines("input.txt")[0]
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func part1(input string) string {
	for i := 1; ; i++ {
		num := strconv.Itoa(i)
		toTest := input + num
		res := md5.Sum([]byte(toTest))
		if res[0] == 0 && res[1] == 0 && (res[2]&0xF0) == 0 {
			return fmt.Sprintf("%s -> %x: %d", toTest, res, i)
		}
	}
}

func part2(input string) string {
	for i := 1; ; i++ {
		num := strconv.Itoa(i)
		toTest := input + num
		res := md5.Sum([]byte(toTest))
		if res[0] == 0 && res[1] == 0 && res[2] == 0 {
			return fmt.Sprintf("%s -> %x: %d", toTest, res, i)
		}
	}
}
