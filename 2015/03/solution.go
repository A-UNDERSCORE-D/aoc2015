package main

import (
	"fmt"
	"strconv"
	"time"

	"awesome-dragon.science/go/adventofcode2015/util"
	"awesome-dragon.science/go/adventofcode2015/util/vector"
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

var dirMap = map[rune]vector.Vec2d{
	'^': vector.V2Up,
	'<': vector.V2Left,
	'>': vector.V2Right,
	'v': vector.V2Down,
}

func part1(input string) string {
	houses := map[vector.Vec2d]int{}
	curPos := vector.Vec2d{X: 0, Y: 0}
	houses[curPos]++

	for _, d := range input {
		curPos = curPos.Add(dirMap[d])
		houses[curPos]++
	}

	return strconv.Itoa(len(houses))
}

func part2(input string) string {
	houses := map[vector.Vec2d]int{}
	curPos := vector.Vec2d{X: 0, Y: 0}
	curRobotPos := vector.Vec2d{X: 0, Y: 0}
	houses[curPos] = 2

	for i, d := range input {
		if i%2 == 0 {
			curPos = curPos.Add(dirMap[d])
			houses[curPos]++
		} else {
			curRobotPos = curRobotPos.Add(dirMap[d])
			houses[curRobotPos]++
		}
	}

	return strconv.Itoa(len(houses))
}
