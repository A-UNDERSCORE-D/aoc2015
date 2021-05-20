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
	var l, w, h int
	total := 0
	for _, line := range input {
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		a1, a2, a3 := l*w, w*h, h*l
		total += util.Sum(2*a1, 2*a2, 2*a3, util.MinOf(a1, a2, a3))
	}
	return strconv.Itoa(total)
}

func part2(input []string) string {
	var l, w, h int
	total := 0
	for _, line := range input {
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		p1, p2, p3 := 2*(l+w), 2*(w+h), 2*(h+l)
		minPerimeter := util.MinOf(p1, p2, p3)
		total += minPerimeter + l*w*h
	}
	return strconv.Itoa(total)
}
