package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strconv"
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

func part1(input []string) string {
	board := make([][]bool, 1000)
	for i := range board {
		board[i] = make([]bool, 1000)
	}

	compiled := regexp.MustCompile(`(?:turn )?(on|off|toggle) (\d+,\d+) through (\d+,\d+)`)

	for _, instruction := range input {
		match := compiled.FindStringSubmatch(instruction)
		ins := match[1]
		start := strings.Split(match[2], ",")
		end := strings.Split(match[3], ",")

		startx := util.MustAtoi(start[0])
		endx := util.MustAtoi(end[0])
		starty := util.MustAtoi(start[1])
		endy := util.MustAtoi(end[1])

		toSetSimple := ins == "on"

		switch ins {
		case "on", "off":
			for x := startx; x <= endx; x++ {
				for y := starty; y <= endy; y++ {
					board[x][y] = toSetSimple
				}
			}

		case "toggle":
			for x := startx; x <= endx; x++ {
				for y := starty; y <= endy; y++ {
					board[x][y] = !board[x][y]
				}
			}
		}

	}
	lit := 0
	// img := image.NewGray(image.Rect(0, 0, 1000, 1000))
	for _, row := range board {
		for _, light := range row {
			if light {
				// img.Set(x, y, color.White)
				lit++
			} /* else {
				// img.Set(x, y, color.Black)
			} */
		}
	}

	// f, err := os.Create("./out.png")
	// if err != nil {
	// 	panic(err)
	// }
	// png.Encode(f, img)
	// fmt.Println(board)
	return strconv.Itoa(lit)
}

func part2(input []string) string {
	board := make([][]int8, 1000)
	for i := range board {
		board[i] = make([]int8, 1000)
	}

	compiled := regexp.MustCompile(`(?:turn )?(on|off|toggle) (\d+,\d+) through (\d+,\d+)`)
	LUT := map[string]int8{
		"on":     +1,
		"off":    -1,
		"toggle": +2,
	}

	for _, instruction := range input {
		match := compiled.FindStringSubmatch(instruction)
		ins := match[1]
		start := strings.Split(match[2], ",")
		end := strings.Split(match[3], ",")

		startx := util.MustAtoi(start[0])
		endx := util.MustAtoi(end[0])
		starty := util.MustAtoi(start[1])
		endy := util.MustAtoi(end[1])

		toAdd := LUT[ins]
		for x := startx; x <= endx; x++ {
			for y := starty; y <= endy; y++ {
				board[x][y] += toAdd
				if board[x][y] < 0 {
					board[x][y] = 0
				}
			}
		}
	}

	total := 0
	img := image.NewGray(image.Rect(0, 0, 1000, 1000))
	for x, r := range board {
		for y, l := range r {
			if l < 0 {
				panic(l)
			}
			total += int(l)
			img.SetGray(x, y, color.Gray{Y: uint8(l)})
		}
	}

	f, err := os.Create("./out2.png")
	if err != nil {
		panic(err)
	}
	png.Encode(f, img)

	return strconv.Itoa(total)
}
