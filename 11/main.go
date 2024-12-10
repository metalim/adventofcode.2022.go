package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	Rounds1 = 20
	Rounds2 = 10000
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
	input = parseInput(string(bs))
	part2(input)
}

type Input []*Monkey

type Monkey struct {
	items       []int
	op          func(int) int
	testDivisor int
	trueMonkey  int
	falseMonkey int
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	catch(err)
	return i
}

func toInts(s []string) []int {
	ints := make([]int, len(s))
	for i, s := range s {
		ints[i] = toInt(s)
	}
	return ints
}

var reMonkey = regexp.MustCompile(`Monkey (\d+):
\s+Starting items: (.*)
\s+Operation: new = old (.) (.*)
\s+Test: divisible by (\d*)
\s+If true: throw to monkey (\d*)
\s+If false: throw to monkey (\d*)`)

func parseInput(input string) Input {
	ms := reMonkey.FindAllStringSubmatch(input, -1)
	monkeys := make(Input, len(ms))
	for _, m := range ms {
		id := toInt(m[1])
		items := toInts(strings.Split(m[2], ", "))
		op := m[3]
		operand := m[4]
		var operandInt int
		if operand != "old" {
			operandInt = toInt(operand)
		}
		divisibleBy := toInt(m[5])
		trueMonkey := toInt(m[6])
		falseMonkey := toInt(m[7])

		monkeys[id] = &Monkey{
			items: items,
			op: func(old int) int {
				if operand == "old" {
					switch op {
					case "+":
						return old + old
					case "*":
						return old * old
					}
				}
				switch op {
				case "+":
					return old + operandInt
				case "*":
					return old * operandInt
				}
				return old
			},
			testDivisor: divisibleBy,
			trueMonkey:  trueMonkey,
			falseMonkey: falseMonkey,
		}
	}
	return monkeys
}

func part1(input Input) {
	timeStart := time.Now()

	inspected := make([]int, len(input))
	for round := 0; round < Rounds1; round++ {
		for id, monkey := range input {
			for i := 0; i < len(monkey.items); i++ {
				inspected[id]++
				level := monkey.items[i]
				level = monkey.op(level)
				level = level / 3
				if level%monkey.testDivisor == 0 {
					input[monkey.trueMonkey].items = append(input[monkey.trueMonkey].items, level)
				} else {
					input[monkey.falseMonkey].items = append(input[monkey.falseMonkey].items, level)
				}
			}
			monkey.items = monkey.items[:0]
		}
	}

	sort.Ints(inspected)
	monkeyBusiness := inspected[len(inspected)-1] * inspected[len(inspected)-2]
	fmt.Printf("Part 1: %d\t\tin %v\n", monkeyBusiness, time.Since(timeStart))
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func part2(input Input) {
	timeStart := time.Now()

	mod := 1
	for _, monkey := range input {
		mod = lcm(mod, monkey.testDivisor)
	}
	fmt.Printf("mod: %d\n", mod)
	inspected := make([]int, len(input))
	for round := 0; round < Rounds2; round++ {
		for id, monkey := range input {
			for i := 0; i < len(monkey.items); i++ {
				inspected[id]++
				level := monkey.items[i]
				level = monkey.op(level)
				level = level % mod // neat trick to keep the numbers small
				if level%monkey.testDivisor == 0 {
					input[monkey.trueMonkey].items = append(input[monkey.trueMonkey].items, level)
				} else {
					input[monkey.falseMonkey].items = append(input[monkey.falseMonkey].items, level)
				}
			}
			monkey.items = monkey.items[:0]
		}
	}

	sort.Ints(inspected)
	monkeyBusiness := inspected[len(inspected)-1] * inspected[len(inspected)-2]
	fmt.Printf("Part 2: %d\t\tin %v\n", monkeyBusiness, time.Since(timeStart))
}
