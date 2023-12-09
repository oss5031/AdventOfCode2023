package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("%s not found\n", os.Args[1])
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	part1, part2 := solve(lines)
	fmt.Printf("%d\n%d\n", part1, part2)

	if err = file.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func solve(lines []string) (part1 int, part2 int) {
	net := make(map[string][]string, len(lines[2:]))
	var start []string
	for _, ln := range lines[2:] {
		net[ln[:3]] = []string{ln[7:10], ln[12:15]}
		if ln[:3][2] == 'A' {
			start = append(start, ln[:3])
		}
	}

	// Part 1
	next := decode(circulator(lines[0]))
	curr := "AAA"
	for curr != "ZZZ" {
		curr = net[curr][next()]
		part1++
	}

	// Part 2
	steps := make([]int, len(start))
	for i, n := range start {
		next = decode(circulator(lines[0]))
		curr = n
		for curr[2] != 'Z' {
			curr = net[curr][next()]
			steps[i]++
		}
	}
	part2 = steps[0]
	for _, s := range steps[1:] {
		part2 = lcm(part2, s)
	}

	return
}

// circulator is a circular iterator of chars
func circulator(x string) func() uint8 {
	i := -1
	return func() uint8 {
		i++
		if i == len(x) {
			i = 0
		}
		return x[i]

	}
}

func decode(nextFunc func() uint8) func() int {
	return func() int {
		switch nextFunc() {
		case 'L':
			return 0
		case 'R':
			return 1
		default:
			return -1
		}
	}
}

// https://en.wikipedia.org/wiki/Euclidean_algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// https://en.wikipedia.org/wiki/Least_common_multiple
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
