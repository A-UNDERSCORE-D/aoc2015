package main

import (
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
	out := 0
	for _, v := range input {
		switch v {
		case '(':
			out += 1

		case ')':
			out -= 1

		}
	}
	return strconv.Itoa(out)
}

func part2(input string) string {
	out := 0
	for i, v := range input {
		switch v {
		case '(':
			out += 1

		case ')':
			out -= 1
		}

		if out < 0 {
			return strconv.Itoa(i + 1)
		}
	}
	return "???"
}
