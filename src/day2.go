package main

import "log"

type RPS int

const (
	Paper RPS = iota
	Rock
	Scissors
)

type match struct{ him, me RPS }

func data2() []match {
	lines := ReadLines("day2.txt")
	var result []match = nil
	for _, line := range lines {
		// split line into opponent and me
		opponent := fromOpponent(line[0:1])
		me := fromMe(line[2:3])
		result = append(result, match{opponent, me})
	}
	return result
}
func dataReco2() []match {
	lines := ReadLines("day2.txt")
	var result []match = nil
	for _, line := range lines {
		// split line into opponent and me
		opponent := fromOpponent(line[0:1])
		me := fromReco(opponent, line[2:3])
		result = append(result, match{opponent, me})
	}
	return result
}

func Day2() {
	log.Print("Running day 2")
	data := data2()

	var globalScore int
	for _, fight := range data {
		globalScore += score(fight)
	}
	// 11150
	log.Printf("Global score from command: %d", globalScore)

	data2 := dataReco2()
	var globalRecoScore int
	for _, fight := range data2 {
		globalRecoScore += score(fight)
	}
	// 8295
	log.Printf("Global score from reco: %d", globalRecoScore)
}

type entry struct {
	him  RPS
	reco string
}

func fromReco(opponent RPS, s string) RPS {
	rules := map[entry]RPS{
		entry{Rock, "X"}:     Scissors, // Lose
		entry{Paper, "X"}:    Rock,
		entry{Scissors, "X"}: Paper,
		entry{Rock, "Y"}:     Rock, // Raw
		entry{Paper, "Y"}:    Paper,
		entry{Scissors, "Y"}: Scissors,
		entry{Rock, "Z"}:     Paper, // Raw
		entry{Paper, "Z"}:    Scissors,
		entry{Scissors, "Z"}: Rock,
	}
	return getOrPanic(rules, entry{opponent, s})
}

func fromOpponent(s string) RPS {
	rules := map[string]RPS{
		"A": Rock,
		"B": Paper,
		"C": Scissors}
	return getOrPanic(rules, s)
}
func fromMe(s string) RPS {
	rules := map[string]RPS{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}
	return getOrPanic(rules, s)
}

func score(fight match) int {
	rules := map[RPS]int{
		Rock:     1,
		Paper:    2,
		Scissors: 3,
	}
	var score int = getOrPanic(rules, fight.me)

	scoreRules := map[match]int{
		match{Scissors, Rock}:     6,
		match{Rock, Paper}:        6,
		match{Paper, Scissors}:    6,
		match{Rock, Rock}:         3,
		match{Paper, Paper}:       3,
		match{Scissors, Scissors}: 3,
	}
	var fightScore = getOrDefault(scoreRules, fight, 0)
	return score + fightScore
}
