package main

import (
	"fmt"
	"strconv"
	"time"

	"awesome-dragon.science/go/adventofcode2015/util"
)

func main() {
	input := util.ReadLines("input.txt")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func part1(input []string) string {
	decodedLen := 0
	encodedLen := 0

	for _, line := range input {
		res, err := strconv.Unquote(line)
		if err != nil {
			fmt.Println(err, line)
			continue
		}

		decodedLen += len(res)
		encodedLen += len(line)
	}
	return fmt.Sprint(encodedLen - decodedLen)
}

func part2(input []string) string {
	reencodedLen := 0
	encodedLen := 0

	for _, line := range input {
		res := strconv.Quote(line)
		reencodedLen += len(res)
		encodedLen += len(line)
	}

	return fmt.Sprint(reencodedLen - encodedLen)
}
