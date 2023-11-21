package main

import (
	"log"
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

func data() map[int]*Elf {
	data := map[int]*Elf{}
	lines := ReadLines("day1.txt")
	var currentElf = &Elf{len(data), []int{}}
	data[len(data)] = currentElf
	for _, line := range lines {
		if line == "" {
			currentElf = &Elf{len(data), []int{}}
			data[len(data)] = currentElf
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
}
