package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s <path> <part>\n", os.Args[0])
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	switch os.Args[2] {
	case "1":
		fmt.Println(part1(string(data)))
	case "2":
		fmt.Println(part2(string(data)))
	default:
		fmt.Printf("usage: %s <path> <part>\n", os.Args[0])
		os.Exit(1)
	}
}

func part1(input string) int {
	var res int
	lines := strings.Split(input, "\n")
	symbols, numbers := parse(lines)
	for p, size := range numbers {
		if hasNeighbours(p.x, p.y, size, &symbols) {
			val, _ := strconv.Atoi(lines[p.y][p.x : p.x+size])
			res += val
		}
	}
	return res
}

func part2(input string) int {
	var res int
	lines := strings.Split(input, "\n")
	symbols, numbers := parse(lines)

	gears := make(map[position]int)
	for p, size := range numbers {
		for _, n := range neighbours(p.x, p.y, size, &symbols) {
			if lines[n.y][n.x] == '*' {
				val, _ := strconv.Atoi(lines[p.y][p.x : p.x+size])
				if _, has := gears[n]; !has {
					gears[n] = val
					continue
				}
				res += gears[n] * val
			}
		}
	}

	return res
}

func neighbours(x, y, size int, symbols *map[position]int) []position {
	var neighs []position
	for _, i := range []int{-1, size} {
		pos := position{x + i, y}
		if _, has := (*symbols)[pos]; has {
			neighs = append(neighs, pos)
		}
	}
	for _, j := range []int{-1, 1} {
		for i := size; i >= -1; i-- {
			pos := position{x + i, y + j}
			if _, has := (*symbols)[pos]; has {
				neighs = append(neighs, pos)
			}
		}
	}
	return neighs
}

func hasNeighbours(x, y, size int, symbols *map[position]int) bool {
	hasXNeighbours := func() bool {
		for _, i := range []int{-1, size} {
			if _, has := (*symbols)[position{x + i, y}]; has {
				return true
			}
		}
		return false
	}
	hasYNeighbours := func() bool {
		for _, j := range []int{-1, 1} {
			for i := size; i >= -1; i-- {
				if _, has := (*symbols)[position{x + i, y + j}]; has {
					return true
				}
			}
		}
		return false
	}
	return hasXNeighbours() || hasYNeighbours()
}

func parse(lines []string) (symbols map[position]int, numbers map[position]int) {
	symbols, numbers = make(map[position]int), make(map[position]int)
	for i, ln := range lines {
		for j := 0; j < len(ln); j++ {
			if ln[j] == '.' {
				continue
			}
			p := position{j, i}
			var size int
			for j+size < len(ln) && ('0' <= ln[j+size] && ln[j+size] <= '9') {
				size++
			}
			if size != 0 {
				numbers[p] = size
				j += size - 1
				continue
			}
			// Everything that is not a number or '.' is a symbol (symbols were not specified)
			symbols[p] = 1
		}
	}
	return // symbols, numbers
}

type position struct {
	x, y int
}
