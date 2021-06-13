package main

import (
	"fmt"
	"strings"
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

func lookAndSay(in string) string {
	out := &strings.Builder{}
	var last rune = 0x0
	count := 0

	for _, r := range in {
		if last == 0x0 {
			// Initialise to -1 to ensure that the below code behaves correctly
			last = r
			// count--
		}

		if r == last {
			count++
		} else {
			fmt.Fprintf(out, "%d%s", count, string(last))
			last = r
			count = 1
		}
	}

	fmt.Fprintf(out, "%d%s", count, string(last))

	return out.String()
}

func part1(input []string) string {
	s := input[0]
	for i := 0; i < 40; i++ {
		new := lookAndSay(s)
		s = new

	}

	return fmt.Sprint(len(s))
}

func part2(input []string) string {
	s := input[0]
	for i := 0; i < 50; i++ {
		new := lookAndSay(s)
		s = new

	}

	return fmt.Sprint(len(s))
}
