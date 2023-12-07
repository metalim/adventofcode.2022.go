package main

import (
	"fmt"
	"os"
	"regexp"
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

	part1(lines)
	part2(lines)
}

func toInts(ss []string) []int {
	ns := make([]int, len(ss))
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		catch(err)
		ns[i] = n
	}
	return ns
}

type Move struct {
	Move, From, To int
}

var reMove = regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

func parse(lines []string) (columns [][]rune, moves []Move) {
	for i, line := range lines {
		if line != "" {
			continue
		}
		n := len(strings.Fields(lines[i-1]))
		columns = make([][]rune, n)
		for j := i - 2; j >= 0; j-- {
			for x := 1; x < len(lines[j]); x += 4 {
				c := rune(lines[j][x])
				if c == ' ' {
					continue
				}
				columns[x/4] = append(columns[x/4], c)
			}
		}

		moves = make([]Move, 0, len(lines)-i-1)
		for _, line := range lines[i+1:] {
			m := reMove.FindStringSubmatch(line)
			if m == nil {
				panic("invalid move: " + line)
			}
			move := toInts(m[1:])
			moves = append(moves, Move{Move: move[0], From: move[1] - 1, To: move[2] - 1})
		}
		return
	}
	panic("no empty line found")
}

func part1(lines []string) {
	columns, moves := parse(lines)
	for _, move := range moves {
		for i := 0; i < move.Move; i++ {
			columns[move.To] = append(columns[move.To], columns[move.From][len(columns[move.From])-1])
			columns[move.From] = columns[move.From][:len(columns[move.From])-1]
		}
	}
	var top []rune
	for _, column := range columns {
		top = append(top, column[len(column)-1])
	}
	fmt.Println("Part 1:", string(top))
}

func part2(lines []string) {
}
