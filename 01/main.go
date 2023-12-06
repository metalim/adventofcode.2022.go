package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go input.txt")
		os.Exit(1)
	}

	bs, err := os.ReadFile(os.Args[1])
	catch(err)
	lines := strings.Split(string(bs), "\n")

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
