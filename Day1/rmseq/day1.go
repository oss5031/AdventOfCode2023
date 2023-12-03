package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	digits = map[string]int{
		"1": 1, "one": 1,
		"2": 2, "two": 2,
		"3": 3, "three": 3,
		"4": 4, "four": 4,
		"5": 5, "five": 5,
		"6": 6, "six": 6,
		"7": 7, "seven": 7,
		"8": 8, "eight": 8,
		"9": 9, "nine": 9,
	}
	trie *Trie
)

func init() {
	trie = New()
	for k, v := range digits {
		if val := trie.Insert(k, v); val != -1 {
			panic("duplicated entry")
		}
	}
}

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
	scanner := bufio.NewScanner(strings.NewReader(input))
	var first, last int32
	for scanner.Scan() {
		first = -1
		for _, c := range scanner.Text() {
			if unicode.IsDigit(c) {
				if first == -1 {
					first = c
				}
				last = c
			}
		}
		// No need to verify if we found something, each line has at least a number
		val, err := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
		if err != nil {
			panic(err)
		}
		res += val
	}
	return res
}

func part2(input string) int {
	var res int
	scanner := bufio.NewScanner(strings.NewReader(input))
	var first, last int
	for scanner.Scan() {
		ln := scanner.Text()
		for i := 0; i < len(ln); i++ {
			if val, has := trie.Search(ln[i:]); has {
				first = val
				break
			}
		}
		for i := len(ln); i >= 0; i-- {
			if val, has := trie.Search(ln[i:]); has {
				last = val
				break
			}
		}
		// No need to verify if we found something, each line has at least number
		val, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
		if err != nil {
			panic(err)
		}
		res += val
	}
	return res
}

func New() *Trie {
	return &Trie{value: -1, children: make(map[int32]*Trie)}
}

type Trie struct {
	value    int
	children map[int32]*Trie
}

func (t *Trie) Insert(key string, val int) int {
	tr := &t
	for _, c := range key {
		ch, has := (*tr).children[c]
		if !has {
			ch = &Trie{value: -1, children: make(map[int32]*Trie)}
			(*tr).children[c] = ch
		}
		tr = &ch
	}
	old := (*tr).value
	(*tr).value = val
	return old
}

func (t *Trie) Search(s string) (int, bool) {
	tr := &t
	for _, c := range s {
		ch, has := (*tr).children[c]
		if !has {
			return -1, false
		}
		tr = &ch
		if (*tr).value != -1 {
			return (*tr).value, true
		}
	}
	return -1, false
}
