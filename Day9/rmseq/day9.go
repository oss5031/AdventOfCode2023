package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func solve(lines []string) (part1, part2 int) {
	for _, ln := range lines {
		var input []int
		for _, f := range strings.Fields(ln) {
			n, _ := strconv.Atoi(f)
			input = append(input, n)
		}

		e, ok := extract(input)
		extracted := [][]int{input}
		// Keep extracting until is no longer possible
		for ok {
			extracted = append(extracted, e)
			e, ok = extract(e)
		}

		var lastRight, lastLeft int
		for i := len(extracted) - 1; i >= 0; i-- {
			currRight, currLeft := extracted[i][len(extracted[i])-1], extracted[i][0]
			lastRight, lastLeft = currRight+lastRight, currLeft-lastLeft
		}
		part1 += lastRight
		part2 += lastLeft
	}
	return
}

func extract(seq []int) (res []int, ok bool) {
	for i := 0; i < len(seq)-1; i++ {
		curr := seq[i+1] - seq[i]
		ok = ok || curr != 0
		res = append(res, curr)
	}
	return
}
