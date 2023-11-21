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
	log.Printf("Global score from command: %d", globalScore)

	data2 := dataReco2()
	var globalRecoScore int
	for _, fight := range data2 {
		globalRecoScore += score(fight)
	}
	log.Printf("Global score from reco: %d", globalRecoScore)
}
func fromReco(opponent RPS, s string) RPS {
	switch s {
	case "X": // Lose
		switch opponent {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	case "Y": // Raw
		switch opponent {
		case Rock:
			return Rock
		case Paper:
			return Paper
		case Scissors:
			return Scissors
		}
	case "Z": // Win
		switch opponent {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	}
	panic("Unknown move")
}
func fromOpponent(s string) RPS {
	switch s {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	default:
		panic("Unknown move")
	}
}
func fromMe(s string) RPS {
	switch s {
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	default:
		panic("Unknown move")
	}
}

func (m RPS) String() string {
	switch m {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		panic("Unknown move")
	}
}
func score(fight match) int {
	var score int
	switch fight.me {
	case Rock:
		score = 1
	case Paper:
		score = 2
	case Scissors:
		score = 3
	default:
		panic("Unknown move")
	}

	var fightScore int
	switch fight {
	case match{Scissors, Rock}:
		fightScore = 6
	case match{Rock, Paper}:
		fightScore = 6
	case match{Paper, Scissors}:
		fightScore = 6
	case match{Rock, Rock}:
		fightScore = 3
	case match{Paper, Paper}:
		fightScore = 3
	case match{Scissors, Scissors}:
		fightScore = 3
	default:
		fightScore = 0
	}

	return score + fightScore
}
