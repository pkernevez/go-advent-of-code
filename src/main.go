package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	var day int
	if len(os.Args) == 1 {
		day = 1
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
	default:
		log.Fatalf("Unsupported day: %d", day)
	}
	log.Print("Finished")

}
