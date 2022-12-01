package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func catch(err error, args ...interface{}) {
	if err == nil {
		return
	}
	if len(args) > 0 {
		panic(errors.Wrap(err, fmt.Sprintf(args[0].(string), args[1:]...)))
	}
	panic(err)
}

func Input() string {
	file := "input.txt"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	input, err := os.ReadFile(file)
	catch(err)
	return strings.TrimSpace(string(input))
}

func Lines(input string) []string {
	return strings.Split(input, "\n")
}

func Ints(lines []string) []int {
	var err error
	ints := make([]int, len(lines))
	for i, line := range lines {
		ints[i], err = strconv.Atoi(line)
		catch(err, "line %d", i)
	}
	return ints
}

func main() {
	lines := Lines(Input())

	// Part 1
	{
		var max, current int
		for _, line := range lines {
			if line == "" {
				if max < current {
					max = current
				}
				current = 0
				continue
			}

			calories, err := strconv.Atoi(line)
			catch(err)

			current += calories
		}
		fmt.Printf("Part 1: %d\n", max)
	}

	// Part 2
	{
		elves := []int{}
		var current int
		for _, line := range lines {
			if line == "" {
				elves = append(elves, current)
				current = 0
				continue
			}

			calories, err := strconv.Atoi(line)
			catch(err)

			current += calories
		}

		sort.IntSlice(elves).Sort()
		top3 := elves[len(elves)-3:]
		sum3 := top3[0] + top3[1] + top3[2]

		fmt.Printf("Part 2: %d\n", sum3)
	}
}
