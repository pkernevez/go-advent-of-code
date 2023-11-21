package main

import (
	"log"
	"sort"
	"strconv"
)

type Elf struct {
	position int
	food     []int
}

func (e *Elf) addFood(food int) {
	e.food = append(e.food, food)
	log.Printf("Elf %d now has %d food", e.position, e.allFood())
}

func (e *Elf) allFood() int {
	sum := 0
	for _, v := range e.food {
		sum += v
	}
	return sum
}

func data() []*Elf {
	var data []*Elf = nil
	lines := ReadLines("day1.txt")
	var currentElf = &Elf{len(data), []int{}}
	data = append(data, currentElf)
	for _, line := range lines {
		if line == "" {
			currentElf = &Elf{len(data), []int{}}
			data = append(data, currentElf)
			log.Print("New Elf")
		} else {
			food, _ := strconv.Atoi(line)
			currentElf.addFood(food)
			log.Printf("Add food %d to Elf %d, contains %d", food, currentElf.position, currentElf.allFood())
		}
	}
	log.Printf("Data contains %d elves", len(data))
	return data
}

func Day1() {
	log.Print("Running day 1")

	data := data()
	// find the elf with the most food
	var maxFood int
	var maxElf int
	for _, elf := range data {
		if elf.allFood() > maxFood {
			maxFood = elf.allFood()
			maxElf = elf.position
		}
	}
	log.Printf("Elf %d has the most food: %d\n", maxElf, maxFood)

	sort.Slice(data, func(i, j int) bool {
		return data[i].allFood() > data[j].allFood()
	})

	sumFoods := 0
	n := 3
	for _, elf := range data[:n] {
		log.Printf("Elf %d has %d food: %v", elf.position, elf.allFood(), elf.food)
		sumFoods += elf.allFood()
	}
	log.Printf("The first %d Elves carry: %d\n", n, sumFoods)
}
