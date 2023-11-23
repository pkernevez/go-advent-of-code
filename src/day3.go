package main

import (
	"fmt"
	"log"
	"strings"
)

func Day3() {
	log.Print("Running day 3")

	data := data3()
	allPrio := 0
	for _, rs := range data {
		log.Printf("priority: %v", rs.priority())
		allPrio += rs.priority()
	}
	log.Printf("total priority: %v", allPrio)
}

type rucksack struct {
	first, second string
}

func data3() []rucksack {
	lines := ReadLines("day3.txt")
	var result []rucksack = nil
	for _, line := range lines {
		first := line[0 : len(line)/2]
		second := line[len(line)/2:]
		result = append(result, rucksack{first, second})
	}
	return result
}

func (rs rucksack) priority() int {
	prio := 0
	for _, c := range rs.sharedItem() {
		prio += priority(c)
	}
	return prio
}

func (rs rucksack) sharedItem() []string {
	var firstChars = strings.Split(rs.first, "")
	var secondChars = strings.Split(rs.second, "")
	withDuplicate := intersection(firstChars, secondChars)
	seen := make(map[string]bool)
	var result []string
	for _, item := range withDuplicate {
		if _, ok := seen[item]; !ok {
			result = append(result, item)
			seen[item] = true
		}
	}
	return result
}

func priority(c string) int {
	b := []byte(c)
	if len(b) != 1 {
		panic(fmt.Sprintf("Char is not only a byte len=%d, %v", len(b), b))
	}
	intValue := b[0]
	if intValue >= 65 && intValue <= 90 {
		return int(intValue) - 65 + 27
	} else if intValue >= 97 && intValue <= 122 {
		return int(intValue) - 97 + 1
	} else {
		panic(fmt.Sprintf("Unknown priority for char [%v], byte is %d, intValue is %d", c, b, intValue))
	}
}
func intersection[T comparable](s1 []T, s2 []T) []T {
	var intersection []T = nil
	for _, first := range s1 {
		for _, second := range s2 {
			if first == second {
				intersection = append(intersection, first)
			}
		}
	}
	return intersection
}
