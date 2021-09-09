package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2015/util"
)

var example = strings.Split(`Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.`, "\n")

func main() {
	RawInput := util.ReadLines("input.txt")
	// RawInput = example
	input := getPeople(RawInput)
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func getPeople(input []string) map[string]map[string]int {
	out := map[string]map[string]int{}
	for _, i := range input {
		split := strings.Split(i, " ")
		name := split[0]
		number, err := strconv.Atoi(split[3])
		if err != nil {
			panic(err)
		}
		otherName := strings.TrimSuffix(split[10], ".")
		modifier := 1
		if split[2] == "lose" {
			modifier = -1
		}

		if out[name] == nil {
			out[name] = make(map[string]int)
		}

		out[name][otherName] = number * modifier
	}

	return out
}

// func recursiveStrat(alreadySeated []string) {
// }
type SeatingPreferences = map[string]map[string]int

func bestSeatingState2(allPrefs SeatingPreferences) map[string][]string {
	names := make([]string, 0, len(allPrefs))
	for n := range allPrefs {
		names = append(names, n)
		// seatingState[n] = make([]string, 0, 2)
	}
	sort.Strings(names)

	states := []map[string][]string{}

	for _, nextName := range names {
		total := 0
		next := nextName
		seatingState := make(map[string][]string)
		for n := range allPrefs {
			seatingState[n] = make([]string, 0, 2)
		}
		for next != "" {
			current := next

			prefs := allPrefs[current]
			best := math.MinInt64
			bestOther := ""
			for other, value := range prefs {
				if util.StringSliceContains(seatingState[other], current) {
					// we cant sit next to someone we're already next to
					continue
				}

				if len(seatingState[other]) >= 2 {
					// we can only have two neighbours for a given person
					continue
				}

				sum := util.Max(value, allPrefs[other][current]) + util.Min(value, allPrefs[other][current])
				if sum > best {
					best = sum
					bestOther = other
				}
			}
			next = bestOther
			if bestOther != "" {
				seatingState[current] = append(seatingState[current], bestOther)
				seatingState[bestOther] = append(seatingState[bestOther], current)
				total += best
			}
		}

		states = append(states, seatingState)
	}

	best := math.MinInt64
	var bestState map[string][]string
outer:
	for _, state := range states {
		for _, n := range state {
			if len(n) != 2 {
				continue outer
			}
		}

		current := 0
		for person, others := range state {
			current += allPrefs[others[0]][person]
			current += allPrefs[others[1]][person]
		}

		if current > best {
			best = current
			bestState = state
		}
	}

	return bestState
}

func part1(input map[string]map[string]int) string {
	best2 := bestSeatingState2(input)
	totalChange2 := 0
	for person, others := range best2 {
		totalChange2 += input[person][others[0]]
		totalChange2 += input[person][others[1]]
	}
	return fmt.Sprint(totalChange2)
}

func part2(input map[string]map[string]int) string {
	me := map[string]int{}

	for name := range input {
		me[name] = 0
		input[name]["me"] = 0
	}

	input["me"] = me

	best2 := bestSeatingState2(input)
	totalChange2 := 0
	for person, others := range best2 {
		totalChange2 += input[person][others[0]]
		totalChange2 += input[person][others[1]]
	}
	return fmt.Sprint(totalChange2)
}
