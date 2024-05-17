package main

import (
	"github.com/eshom/aoc2022-go/pkg/assert"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// part 1
	var input string = readInput("cmd/day1/data/input.txt")
	var calories []string = strings.Split(input, "\n")
	var sums []int = sumCals(calories)
	var maxCal int = maxSum(sums)
	log.Println(maxCal)

	// part 2
	var top3 []int = sums
	slices.Sort(top3)
	slices.Reverse(top3)
	assert.Assert(len(top3) >= 3, "must have at least 3 sums")
	top3 = top3[0:3]
	log.Println(top3[0] + top3[1] + top3[2])
}

func readInput(filename string) string {
	contents, err := os.ReadFile(filename)
	assert.NoError(err)
	return string(contents)
}

func sumCals(cals []string) []int {
	out := make([]int, 0, 1000)
	var sum int = 0
	for _, str := range cals {
		if str != "\n" && str != "" {
			num, err := strconv.Atoi(str)
			assert.NoError(err)
			sum += num
		} else {
			out = append(out, sum)
			sum = 0
		}
	}

	return out
}

func maxSum(sums []int) int {
	var out int = 0
	for _, s := range sums {
		out = max(out, s)
	}
	return out
}
