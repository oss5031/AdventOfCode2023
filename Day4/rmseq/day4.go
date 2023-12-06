package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s <path> <part>\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	switch os.Args[2] {
	case "1":
		fmt.Println(part1(lines))
	case "2":
		fmt.Println(part2(lines))
	default:
		fmt.Printf("usage: %s <path> <part>\n", os.Args[0])
		os.Exit(1)
	}
}

func part1(lines []string) int {
	var res int
	for _, ln := range lines {
		if m := matches(fields(ln)); m >= 1 {
			res += m << (m - 1)
		}
	}
	return res
}

func part2(lines []string) int {
	cards, total := make([]int, len(lines)), 0
	for i := 0; i < len(lines); i++ {
		cards[i]++
		total += cards[i]
		for j := 1; j <= matches(fields(lines[i])); j++ {
			cards[i+j] += cards[i]
		}
	}
	return total
}

func fields(ln string) ([]string, []string) {
	ln = ln[strings.Index(ln, ":")+1:]
	sp := strings.Index(ln, "|")
	return strings.Fields(ln[:sp]), strings.Fields(ln[sp:])
}

func matches(x, y []string) int {
	set := make(map[string]bool)
	for _, s := range y {
		set[s] = true
	}
	var res int
	for _, s := range x {
		if _, has := set[s]; has {
			res++
		}
	}
	return res
}
