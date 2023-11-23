package main

import (
	"log"
	"os"
	"strconv"
)

// Main REX
// Pas d'operateur Ternaire
// Pas de return value dans if or switch
//  pas de déconstructeur
// Impossible de compiler avec une varaible non utilisée...
// Pas de char

func main() {
	var day int
	if len(os.Args) == 1 {
		day = 3
		log.Printf("Use default day: %v\n", day)
	} else {
		explicitDay, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("error: %v\n", err)
		}
		day = explicitDay
	}
	log.Printf("Running day %d\n", day)

	switch day {
	case 1:
		Day1()
	case 2:
		Day2()
	case 3:
		Day3()
	default:
		log.Fatalf("Unsupported day: %d", day)
	}
	log.Print("Finished")

}
