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

	dict  *Dict
	input string
)

func init() {
	data, err := os.ReadFile("resources/day1.txt")
	if err != nil {
		panic(err)
	}
	input = string(data)

	dict = New()
	for k, v := range digits {
		if val := dict.Insert(k, v); val != -1 {
			panic("duplicated value")
		}
	}
}

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
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
			_, val, has := dict.Search(ln[i:])
			if has {
				first = val
				break
			}
		}
		for i := len(ln); i >= 0; i-- {
			_, val, has := dict.Search(ln[i:])
			if has {
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

func New() *Dict {
	return &Dict{Key: -1, Val: -1}
}

// Dict is definitely not a tree
type Dict struct {
	Key      int32
	Val      int
	Children []*Dict // TODO: Should it be a map?
}

func (d *Dict) Insert(key string, val int) int {
	tr := &d
	for _, c := range key {
		ch := (*tr).child(c)
		if ch == nil {
			ch = &Dict{Key: c, Val: -1}
			(*tr).Children = append((*tr).Children, ch)
		}
		tr = &ch
	}
	old := (*tr).Val
	(*tr).Val = val
	return old
}

func (d *Dict) Search(s string) (string, int, bool) {
	var match string
	tr := &d
	for _, c := range s {
		ch := (*tr).child(c)
		if ch == nil {
			return "", -1, false
		}
		tr = &ch
		match += string(c)
		if (*tr).Val != -1 {
			return match, (*tr).Val, true
		}
	}
	return "", -1, false
}

func (d *Dict) child(c int32) *Dict {
	for _, ch := range d.Children {
		if ch.Key == c {
			return ch
		}
	}
	return nil
}
