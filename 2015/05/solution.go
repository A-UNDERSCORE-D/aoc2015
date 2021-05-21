package main

import (
	"fmt"
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
	isNiceString := func(s string) bool {
		vowelCount := 0
		doubleCount := 0
		if util.ContainsAny(s, "ab", "cd", "pq", "xy") {
			return false
		}

		var prevRune rune = 0
		for _, r := range s {
			switch r {
			case 'a', 'e', 'i', 'o', 'u':
				vowelCount++
			}
			if r == prevRune {
				doubleCount++
			}
			prevRune = r
		}

		return vowelCount >= 3 && doubleCount > 0
	}

	return fmt.Sprint(len(util.FilterStrSlice(input, isNiceString)))
}

func slidingWindowStr(size int, str string) []string {
	out := make([]string, 0)

	for i := 0; i < len(str); i++ {
		window := str[i:util.Min(len(str), i+size)]
		out = append(out, window)
	}

	return out
}

func part2(input []string) string {
	isNiceString2 := func(s string) bool {
		pairs := map[string]int{}
		hasRepeatWithSingleLetterSpace := false
		previousWindow := ""
		for _, w := range slidingWindowStr(2, s) {
			if len(w) == 2 && w != previousWindow {
				pairs[w]++
			}
			previousWindow = w
		}

		for _, w := range slidingWindowStr(3, s) {
			if len(w) != 3 {
				continue
			}
			if w[0] == w[2] && w[1] != w[0] {
				hasRepeatWithSingleLetterSpace = true
			}
		}

		for _, i := range pairs {
			if i == 2 {
				return hasRepeatWithSingleLetterSpace
			}
		}
		return false
	}

	return fmt.Sprint(len(util.FilterStrSlice(input, isNiceString2)))
}
