package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"awesome-dragon.science/go/adventofcode2015/util"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func findAllNumbers(v interface{}, ignoreRed bool) []float64 {
	out := []float64{}

	switch real := v.(type) {
	case float64:
		out = append(out, real)

	case map[string]interface{}:
		if ignoreRed {
			for key, value := range real {
				if key == "red" {
					return nil
				}

				if res, ok := value.(string); ok && res == "red" {
					return nil
				}
			}
		}
		for _, item := range real {
			out = append(out, findAllNumbers(item, ignoreRed)...)
		}
	case []interface{}:
		for _, item := range real {
			out = append(out, findAllNumbers(item, ignoreRed)...)
		}

		// default:
		// 	fmt.Printf("Dont know what to do with type %t! (%#v)\n", real, real)

	}
	return out
}

func part1(input []byte) string {
	var data interface{}
	json.Unmarshal(input, &data)
	// fmt.Println(data)
	// fmt.Println(findAllNumbers(data))

	return fmt.Sprint(util.SumFloats(findAllNumbers(data, false)...))
}

func part2(input []byte) string {
	var data interface{}
	if err := json.Unmarshal(input, &data); err != nil {
		fmt.Println(err)
	}
	return fmt.Sprint(util.SumFloats(findAllNumbers(data, true)...))
}
