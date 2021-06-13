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
	res = part2(res)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

const alphabet = `abcdefghijklmnopqrstuvwxyz`

func alphaIdx(r byte) int {
	return int(r - 'a')
}

func incLetter(r byte) (out byte, rollover bool) {
	idx := alphaIdx(r)
	inced := idx + 1
	if inced >= len(alphabet) {
		return 'a', true
	}

	return alphabet[inced], false
}

func incString(in []byte) []byte {
	out := make([]byte, len(in))
	copy(out, in)
	for i := len(in) - 1; i >= 0; i-- {
		new, rolled := incLetter(in[i])
		out[i] = new

		if !rolled {
			break
		}
	}

	return out
}

func slidingWindowStr(size int, str string) []string {
	out := make([]string, 0)

	for i := 0; i < len(str); i++ {
		window := str[i:util.Min(len(str), i+size)]
		out = append(out, window)
	}

	return out
}

func passwordIsValid(p string) bool {
	if strings.ContainsAny(p, "iol") {
		return false
	}

	hasTriplet := false
	for _, triplet := range slidingWindowStr(3, p) {
		// Likely breaks at the end.
		if len(triplet) < 3 {
			continue
		}
		start := triplet[0]
		mid := triplet[1]
		end := triplet[2]
		iStart, rollover1 := incLetter(start)
		iMid, rollover2 := incLetter(mid)

		if rollover1 || rollover2 {
			continue
		}

		if iStart != mid || iMid != end {
			continue
		}

		hasTriplet = true
		break
	}

	if !hasTriplet {
		return false
	}

	lastPairIdx := -1
	pairCount := 0
	for i, pair := range slidingWindowStr(2, p) {
		if pairCount >= 2 {
			break
		}
		if len(pair) < 2 {
			continue
		}

		isPair := pair[0] == pair[1]
		if isPair && i != lastPairIdx+1 {
			pairCount++
			lastPairIdx = i
		}
	}

	return pairCount >= 2
}

func nextValidPassword(start string) string {
	buf := make([]byte, len(start))
	copy(buf, start)
	for {
		if passwordIsValid(string(buf)) {
			return string(buf)
		}
		buf = incString(buf)
	}
}

func part1(input []string) string {
	return nextValidPassword(input[0])
}

func part2(input string) string {
	res := nextValidPassword(string(incString([]byte(input))))
	fmt.Println(res, passwordIsValid(res))
	return ""
}
