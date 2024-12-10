package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: go run main.go input.txt")
		os.Exit(1)
	}

	bs, err := os.ReadFile(flag.Arg(0))
	catch(err)

	input := parseInput(string(bs))
	part1(input)
	part2(input)
}

type List []any
type Input []List

var reNum = regexp.MustCompile(`\d+`)

func parseList(input string) List {
	stack := []List{nil}
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '[':
			stack = append(stack, List{})
		case ']':
			l := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = append(stack[len(stack)-1], l)
		case ',':
			continue
		default:
			m := reNum.FindStringIndex(input[i:])
			if m == nil {
				panic(fmt.Sprintf("invalid list: %s at %d", input, i))
			}
			n, err := strconv.Atoi(input[i+m[0] : i+m[1]])
			catch(err)
			stack[len(stack)-1] = append(stack[len(stack)-1], n)
			i += m[1] - 1
		}
	}
	return stack[0][0].(List)
}

func parseInput(input string) Input {
	lines := strings.Split(input, "\n")
	packets := make(Input, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		packets = append(packets, parseList(line))
	}
	return packets
}

func compare(a, b List) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		var v int
		switch aa := a[i].(type) {
		case int:
			switch bb := b[i].(type) {
			case int:
				v = aa - bb
			case List:
				v = compare(List{aa}, bb)
			}
		case List:
			switch bb := b[i].(type) {
			case int:
				v = compare(aa, List{bb})
			case List:
				v = compare(aa, bb)
			}
		}
		if v != 0 {
			return v
		}
	}
	return len(a) - len(b)
}

func part1(input Input) {
	timeStart := time.Now()
	sum := 0
	for i := 0; i < len(input); i += 2 {
		if compare(input[i], input[i+1]) < 0 {
			sum += i/2 + 1
		}
	}

	fmt.Printf("Part 1: %d\t\tin %v\n", sum, time.Since(timeStart))
}

func part2(input Input) {
	timeStart := time.Now()
	p2 := List{List{2}}
	p6 := List{List{6}}
	packets := append(input, p2, p6)
	slices.SortFunc(packets, compare)
	p2Index, _ := slices.BinarySearchFunc(packets, p2, compare)
	p6Index, _ := slices.BinarySearchFunc(packets, p6, compare)
	decoderKey := (p2Index + 1) * (p6Index + 1)

	fmt.Printf("Part 2: %d\t\tin %v\n", decoderKey, time.Since(timeStart))
}
